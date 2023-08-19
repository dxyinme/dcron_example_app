package controllers

import "github.com/gin-gonic/gin"

type TasksController struct{}

//	List godoc
//
//	@Summary			list tasks
//	@Description	list cron tasks
//	@Tags					tasks
//	@Accept				json
//	@Produce			json
//	@Success			200 {object} []types.Task
//	@Router				/tasks [get]
func (tc *TasksController) List(ctx *gin.Context) {

}
