package controllers

import (
	"app/controllers/helper"
	"app/internal/db"
	"app/internal/logic"
	"app/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TaskController struct {
	taskl logic.TaskLogic
}

//	CreateOrUpdate godoc
//
//	@Summary			create or update task
//	@Description	create or update cron task
//	@Tags					task
//	@Accept				json
//	@Produce			json
//	@Param				taskName path string true "cron task name"
//	@Param				body body types.TaskReq true "create or update cron task request"
//	@Success			200 {string} OK
//	@Failure			400 {string} BadRequest
//	@Router				/tasks/{taskName} [put]
func (tc *TaskController) CreateOrUpdate(ctx *gin.Context) {
	var (
		taskName string = helper.RemoveSlash(ctx.Param("taskName"))
		err      error
		taskReq  = types.TaskReq{}
	)

	if err = ctx.ShouldBindJSON(&taskReq); err != nil {
		logrus.Error(err)
		ctx.String(http.StatusBadRequest, "body not bind to types.TaskReq")
		return
	}
	taskData := db.Task{
		Name:           taskName,
		CronStr:        taskReq.CronStr,
		DBCustomerName: taskReq.DBName,
		SQLStr:         taskReq.SQLStr,
	}

	if err = tc.taskl.UpsertCronTask(&taskData); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err = tc.taskl.UpsertCronTaskToDcron(&taskData); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.String(http.StatusOK, "OK")
}

// Remove godoc
//
//	@Summary			remove task
//	@Description	remove cron task by task name
//	@Tags					task
//	@Accept				json
//	@Produce			json
//	@Param				taskName path string true "cron task name"
//	@Success			200	{string} OK
//	@Failure			400 {string} BadRequest
//	@Router				/tasks/{taskName} [delete]
func (tc *TaskController) Remove(ctx *gin.Context) {
	var (
		taskName string = helper.RemoveSlash(ctx.Param("taskName"))
		err      error
	)

	tc.taskl.RemoveCronTaskFromDcron(taskName)
	if err = tc.taskl.DeleteCronTask(taskName); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK, "OK")
}

//	List godoc
//
//	@Summary			list tasks
//	@Description	list cron tasks
//	@Tags					task
//	@Accept				json
//	@Produce			json
//	@Success			200 {object} []types.Task
//	@Router				/tasks [get]
func (tc *TaskController) List(ctx *gin.Context) {
	tasksResponse := make([]types.Task, 0)
	ss := db.SelfStoreUtil{}.I()
	limit := 10
	lastId := uint(0)
	for {
		tasks, err := ss.GetTasksByIDLimit(lastId, limit)
		if err != nil {
			logrus.Error(err)
			continue
		}
		for _, task := range tasks {
			webTask := types.Task{}
			webTask.FromDBTask(&task)
			tasksResponse = append(tasksResponse, webTask)
		}
		ltasks := len(tasks)
		if ltasks < limit {
			break
		}
		lastId = tasks[ltasks-1].ID + 1
	}
	ctx.JSON(http.StatusOK, tasksResponse)
}

//	Task Get godoc
//
//	@Summary			task Get
//	@Description  task Get
//	@Tags					task
//	@Accept				json
//	@Produce			json
//	@Param				taskName path string true "task name"
//	@Success			200 {object} types.Task
//	@Router				/tasks/{taskName} [get]
func (tc *TaskController) Get(ctx *gin.Context) {
	var (
		taskName string = helper.RemoveSlash(ctx.Param("taskName"))
		response types.Task
		err      error
		taskData db.Task
	)

	if taskData, err = tc.taskl.GetCronTask(taskName); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	response.FromDBTask(&taskData)
	ctx.JSON(http.StatusOK, &response)
}
