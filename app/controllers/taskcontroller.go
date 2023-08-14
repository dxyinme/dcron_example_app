package controllers

import "github.com/gin-gonic/gin"

type TaskController struct{}

//	Add godoc
//
//	@Summary			add task
//	@Description	add cron task
//	@Tags					task
//	@Accept				json
//	@Produce			json
//	@Param				taskName path string true "cron task name"
//	@Success			200 {string} OK
//	@Failure			400 {string} BadRequest
//	@Router				/task/{taskName} [post]
func (tc *TaskController) Add(ctx *gin.Context) {}

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
