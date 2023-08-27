package controllers

import "github.com/gin-gonic/gin"

type TaskController struct{}

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
func (tc *TaskController) CreateOrUpdate(ctx *gin.Context) {}

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
func (tc *TaskController) Remove(ctx *gin.Context) {}

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
func (tc *TaskController) Get(ctx *gin.Context) {}
