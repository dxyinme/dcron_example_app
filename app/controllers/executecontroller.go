package controllers

import (
	"app/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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
//	@Router				/execute/runSQL [get]
func (ec *ExecuteController) RunSQL(ctx *gin.Context) {
	var err error
	req := types.RunSQLReq{}
	err = ctx.BindJSON(&req)
	if err != nil {
		logrus.Error(err)
		ctx.String(http.StatusBadRequest, "binding error: %s", err.Error())
		return
	}

	ctx.String(http.StatusOK, "OK")
}
