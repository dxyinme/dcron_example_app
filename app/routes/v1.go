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
	tasksController := new(controllers.TasksController)
	executeController := new(controllers.ExecuteController)
	dbController := new(controllers.DBController)

	v1 := eng.Group(basePath)
	{
		task := v1.Group("/task")
		{
			task.PUT("/:taskName", taskController.CreateOrUpdate)
			task.DELETE("/:taskName", taskController.Remove)
		}

		v1.GET("/tasks", tasksController.List)

		execute := eng.Group("/execute")
		{
			execute.POST("/runSQL", executeController.RunSQL)
		}

		databases := eng.Group("/databases")
		{
			databases.PUT("/:dbName", dbController.CreateOrUpdate)
		}
	}
	return
}
