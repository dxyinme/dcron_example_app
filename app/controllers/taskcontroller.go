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
		taskName  string
		uriHelper *helper.UriHelper = nil
		exist     bool              = false
		err       error
		taskReq   = types.TaskReq{}
	)
	if uriHelper, err = helper.GetUriHelperFromGinCtx(ctx); err != nil {
		logrus.Error(err)
		return
	}
	if taskName, exist = uriHelper.Get("taskName"); !exist {
		logrus.Error("taskName not exist in uri")
		ctx.String(http.StatusBadRequest, "taskName not exist in uri")
		return
	}

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
		taskName  string
		uriHelper *helper.UriHelper = nil
		exist     bool              = false
		err       error
	)
	if uriHelper, err = helper.GetUriHelperFromGinCtx(ctx); err != nil {
		logrus.Error(err)
		return
	}
	if taskName, exist = uriHelper.Get("taskName"); !exist {
		logrus.Error("taskName not exist in uri")
		ctx.String(http.StatusBadRequest, "taskName not exist in uri")
		return
	}

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
func (tc *TaskController) List(ctx *gin.Context) {}

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
		taskName  string
		uriHelper *helper.UriHelper = nil
		exist     bool              = false
		response  types.Task
		err       error
		taskData  db.Task
	)
	if uriHelper, err = helper.GetUriHelperFromGinCtx(ctx); err != nil {
		logrus.Error(err)
		return
	}
	if taskName, exist = uriHelper.Get("taskName"); !exist {
		logrus.Error("taskName not exist in uri")
		ctx.String(http.StatusBadRequest, "taskName not exist in uri")
		return
	}
	if taskData, err = tc.taskl.GetCronTask(taskName); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	response.FromDBTask(&taskData)
	ctx.JSON(http.StatusOK, &response)
}
