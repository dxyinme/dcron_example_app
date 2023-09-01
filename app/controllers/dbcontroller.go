package controllers

import (
	"app/controllers/helper"
	"app/internal/customerdb"
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
//	@Param				dbName path string true "database customer name"
//	@Param				body body types.DBReq true "database CreateOrUpdate request"
//	@Success			200 {string} OK
//	@Router				/databases/{dbName} [put]
func (c *DBController) CreateOrUpdate(ctx *gin.Context) {
	var (
		err       error = nil
		dbName    string
		exist     bool
		dbReq                       = types.DBReq{}
		uriHelper *helper.UriHelper = nil
	)

	if uriHelper, err = helper.GetUriHelperFromGinCtx(ctx); err != nil {
		logrus.Error(err)
		return
	}

	dbName, exist = uriHelper.Get("dbName")
	if !exist {
		logrus.Error("dbName not exist in uri")
		ctx.String(http.StatusBadRequest, "dbName not exist in uri")
		return
	}

	if err = ctx.ShouldBindJSON(&dbReq); err != nil {
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

	err = customerdb.DBStoresUtil{}.I().Add(dbData)
	if err != nil {
		logrus.Error(err)
		ctx.String(http.StatusInternalServerError, err.Error())
		if err = ss.DeleteDataBaseByCustomerName(dbName); err != nil {
			logrus.Error(err)
		}
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
//	@Param				dbName path string true "database customer name"
//	@Success			200 {string} OK
//	@Router				/databases/{dbName} [delete]
func (c *DBController) Remove(ctx *gin.Context) {
	var (
		err       error
		dbName    string
		exist     bool              = false
		uriHelper *helper.UriHelper = nil
	)
	if uriHelper, err = helper.GetUriHelperFromGinCtx(ctx); err != nil {
		logrus.Error(err)
		return
	}
	if dbName, exist = uriHelper.Get("dbName"); !exist {
		logrus.Error("dbName not exist in uri")
		ctx.String(http.StatusBadRequest, "dbName not exist in uri")
		return
	}
	ss := db.SelfStoreUtil{}.I()
	if err = ss.DeleteDataBaseByCustomerName(dbName); err != nil {
		logrus.Error(err.Error())
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK, "OK")
}

//	DB Get godoc
//
//	@Summary			db Get
//	@Description  db Get
//	@Tags					db
//	@Accept				json
//	@Produce			json
//	@Param				dbName path string true "database customer name"
//	@Success			200 {object} types.DB
//	@Router				/databases/{dbName} [get]
func (c *DBController) Get(ctx *gin.Context) {
	var (
		err       error
		dbName    string
		dbData    = db.Database{}
		respDB    = types.DB{}
		exist     bool
		uriHelper *helper.UriHelper = nil
	)

	if uriHelper, err = helper.GetUriHelperFromGinCtx(ctx); err != nil {
		logrus.Error(err)
		return
	}

	dbName, exist = uriHelper.Get("dbName")
	if !exist {
		logrus.Error("dbName not exist in uri")
		ctx.String(http.StatusBadRequest, "dbName not exist in uri")
		return
	}

	ss := db.SelfStoreUtil{}.I()
	if dbData, err = ss.GetDataBaseByCustomerName(dbName); err != nil {
		logrus.Error(err)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	respDB.FromDBDatabase(&dbData)
	ctx.JSON(http.StatusOK, &respDB)
}

//	DB List godoc
//
//	@Summary			list Database
//	@Description	list database
//	@Tags					db
//	@Accept				json
//	@Produce			json
//	@Success			200 {object} []types.DB
//	@Router				/databases [get]
func (c *DBController) List(ctx *gin.Context) {}
