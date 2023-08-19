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
//	@Router				/task/{taskName} [post]
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
//	@Router				/task/{taskName} [delete]
func (tc *TaskController) Remove(ctx *gin.Context) {}
