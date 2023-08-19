package controllers

import "github.com/gin-gonic/gin"

type ExecuteController struct{}

//	RunSQL godoc
//
//	@Summary			execute sql cmd
//	@Description	execute sql cmd
//	@Tags					execute
//	@Accept				json
//	@Produce			json
//	@Param				body body types.RunSQLReq true "Run SQL request"
//	@Success			200 {string} OK
//	@Router				/execute [get]
func (ec *ExecuteController) RunSQL(ctx *gin.Context) {}
