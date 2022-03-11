package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTasksList(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	//从gin.Keys中取服务
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	taskResp, err := taskService.GetTasksList(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"task":  taskResp.TaskList,
			"count": taskResp.Count,
		},
	})
}

func CreateTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	//从gin.Keys中取得服务实例
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	taskRes, err := taskService.CreateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": taskRes.TaskDetail,
	})
}

func GetTaskDetail(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	taskResp, err := taskService.GetTask(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": taskResp.TaskDetail,
	})
}

func DeleteTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)
	taskResp, err := taskService.DeleteTask(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": taskResp.TaskDetail,
	})
}

func UpdateTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)
	taskResp, err := taskService.UpdateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": taskResp.TaskDetail,
	})
}
