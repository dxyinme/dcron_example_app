package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

// it is not necessary to generate the
// internal call swagger file.
func InjectInternalCall(eng *gin.Engine) {
	basePath := "/internalCall"

	internalController := new(controllers.InternalController)
	internalCall := eng.Group(basePath)
	{
		internalCall.POST("/addTask", internalController.AddTask)
		internalCall.POST("/removeTask", internalController.RemoveTask)

		internalCall.POST("/addDatabase")
		internalCall.POST("/removeDatabase")

		internalCall.POST("/checkAlive", internalController.CheckAlive)
	}
}
