package controllers

import (
	"app/internal/db"
	"app/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DBController struct {
}

//	DB CreateOrUpdate godoc
//
//	@Summary			db CreateOrUpdate
//	@Description  db CreateOrUpdate
//	@Tags					db
//	@Accept				json
//	@Produce			json
//	@Param				dbName path string true "database custom name"
//	@Param				body body types.DBReq true "database CreateOrUpdate request"
//	@Success			200 {string} OK
//	@Router				/databases/{dbName} [put]
func (c *DBController) CreateOrUpdate(ctx *gin.Context) {
	var (
		err    error = nil
		dbName string
		dbReq  = types.DBReq{}
	)
	err = ctx.ShouldBindUri(&dbName)
	if err != nil {
		logrus.Error(err)
		return
	}
	err = ctx.ShouldBindJSON(&dbReq)
	if err != nil {
		logrus.Error(err)
		return
	}
	ss := db.SelfStoreUtil{}.I()
	dbData := &db.Database{
		CustomerName: dbName,
		User:         dbReq.User,
		DBType:       dbReq.DBType,
		Password:     dbReq.Password,
		DatabaseName: dbReq.DatabaseName,
		Addr:         dbReq.Addr,
	}
	err = ss.UpsertDataBase(dbData)
	if err != nil {
		logrus.Error(err)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = db.DBStoresUtil{}.I().Add(dbData)
	if err != nil {
		logrus.Error(err)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK, "OK")
}

//	DB Remove godoc
//
//	@Summary			db Remove
//	@Description  db Remove
//	@Tags					db
//	@Accept				json
//	@Produce			json
//	@Param				dbName path string true "database custom name"
//	@Success			200 {string} OK
//	@Router				/databases/{dbName} [delete]
func (c *DBController) Remove(ctx *gin.Context) {}
