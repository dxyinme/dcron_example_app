package controllers

import (
	"app/internal/db"
	"app/internal/logic"
	"app/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type InternalController struct {
	dbl   logic.DBLogic
	taskl logic.TaskLogic
}

func (ic InternalController) CheckAlive(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

func (ic InternalController) AddTask(ctx *gin.Context) {
	var (
		err  error
		task types.Task
	)
	if err = ctx.ShouldBindJSON(&task); err != nil {
		logrus.Error(err)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	if err = ic.taskl.UpsertCronTaskToDcron(&db.Task{
		Name:           task.Name,
		CronStr:        task.CronStr,
		SQLStr:         task.SQLStr,
		DBCustomerName: task.DBName,
	}); err != nil {
		logrus.Error(err)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK, "OK")
}

func (ic InternalController) RemoveTask(ctx *gin.Context) {
	var (
		err error
		req types.RemoveTaskReq
	)
	if err = ctx.ShouldBindJSON(&req); err != nil {
		logrus.Error(err)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ic.taskl.RemoveCronTaskFromDcron(req.Name)
	ctx.String(http.StatusOK, "OK")
}
