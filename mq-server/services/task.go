package services

import (
	"encoding/json"
	"log"
	"mq-server/model"
)

//从Rabbit中接受信息，写入数据库
func CreateTask() {
	ch, err := model.MQ.Channel()
	if err != nil {
		panic(err)
	}
	q, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	err = ch.Qos(1, 0, false)
	if err != nil {
		panic(err)
	}
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	//处于监听状态，监听生产端的生产，要阻塞主进程
	go func() {
		for d := range msgs {
			var t model.Task
			err := json.Unmarshal(d.Body, &t)
			if err != nil {
				panic(err)
			}
			model.DB.Create(&t)
			log.Println("Done")
			_ = d.Ack(false)
		}
	}()
}
