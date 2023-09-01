package controllers

import (
	"app/internal/customerdb"
	"app/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
//	@Router				/execute/runSQL [post]
func (ec *ExecuteController) RunSQL(ctx *gin.Context) {
	var (
		err    error
		exist  bool
		dbconn *gorm.DB = nil
	)
	req := types.RunSQLReq{}
	if err = ctx.BindJSON(&req); err != nil {
		logrus.Error(err)
		ctx.String(http.StatusBadRequest, "binding error: %s", err.Error())
		return
	}

	dbs := customerdb.DBStoresUtil{}.I()
	if dbconn, exist = dbs.Load(req.DBCustomerName); !exist {
		err = fmt.Errorf("database `%s` not exist", req.DBCustomerName)
		logrus.Error(err)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err = dbconn.Exec(req.SQLStr).Error; err != nil {
		logrus.Error(err)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.String(http.StatusOK, "OK")
}
