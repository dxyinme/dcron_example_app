package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func New() (eng *gin.Engine) {
	basePath := "/api/v1"
	eng = gin.Default()
	eng.Use(
		gin.Recovery(),
	)
	taskController := new(controllers.TaskController)
	v1 := eng.Group(basePath)
	{
		task := v1.Group("/task")
		{
			task.POST("/:taskName", taskController.Add)
			task.DELETE("/:taskName", taskController.Remove)
		}
	}
	return
}
