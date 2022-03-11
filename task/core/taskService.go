package core

import (
	"context"
	"encoding/json"
	"errors"
	"task/model"
	"task/services"

	"github.com/streadway/amqp"
)

//创建备忘录，放到rabbitMQ队列中
func (*TaskService) CreateTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	ch, err := model.MQ.Channel()
	if err != nil {
		err = errors.New("rabbitMQ channel err:" + err.Error())
		return err
	}
	q, err := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	if err != nil {
		err = errors.New("rabbitMQ queue err:" + err.Error())
		return err

	}
	body, _ := json.Marshal(req)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		err = errors.New("rabbitMQ publish err:" + err.Error())
		return err

	}
	return nil
}

//实现服务接口，获取备忘录列表
func (*TaskService) GetTasksList(ctx context.Context, req *services.TaskRequest, resp *services.TaskListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	var taskData []model.Task
	var count uint32
	//查找
	err := model.DB.Offset(req.Start).Limit(req.Limit).Where("uid=?", req.Uid).Find(&taskData).Error
	if err != nil {
		return errors.New("mysql find:" + err.Error())
	}
	//统计数量
	model.DB.Model(&model.Task{}).Where("uid=?", req.Uid).Count(&count)
	var taskRes []*services.TaskModel
	for _, item := range taskData {
		taskRes = append(taskRes, BuildTask(item))
	}
	resp.TaskList = taskRes
	resp.Count = count
	return nil
}

//获取详细备忘录
func (*TaskService) GetTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	taskData := model.Task{}
	model.DB.First(&taskData, req.Id)
	taskRes := BuildTask(taskData)
	resp.TaskDetail = taskRes
	return nil
}

//修改备忘录
func (*TaskService) UpdateTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	taskData := model.Task{}
	model.DB.Model(&model.Task{}).Where("id=? AND uid=?", req.Id, req.Uid).First(&taskData)
	taskData.Title = req.Title
	taskData.Content = req.Content
	taskData.Status = int(req.Status)
	model.DB.Save(&taskData)
	resp.TaskDetail = BuildTask(taskData)
	return nil
}

//删除备忘录
func (*TaskService) DeleteTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	err := model.DB.Model(&model.Task{}).Where("id=? AND uid=?", req.Id, req.Uid).Delete(&model.Task{}).Error
	if err != nil {
		return errors.New("delete err:" + err.Error())
	}
	return nil
}
