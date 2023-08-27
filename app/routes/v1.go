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
	executeController := new(controllers.ExecuteController)
	dbController := new(controllers.DBController)

	v1 := eng.Group(basePath)
	{
		task := v1.Group("/tasks")
		{
			task.PUT("/:taskName", taskController.CreateOrUpdate)
			task.DELETE("/:taskName", taskController.Remove)
			task.GET("/:taskName", taskController.Get)
		}

		v1.GET("/tasks", taskController.List)

		execute := v1.Group("/execute")
		{
			execute.POST("/runSQL", executeController.RunSQL)
		}

		databases := v1.Group("/databases", dbController.List)
		{
			databases.PUT("/:dbName", dbController.CreateOrUpdate)
			databases.DELETE("/:dbName", dbController.Remove)
			databases.GET("/:dbName", dbController.Get)
		}
		v1.GET("/databases", dbController.List)
	}
	return
}
