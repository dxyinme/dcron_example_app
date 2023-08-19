package controllers

import "github.com/gin-gonic/gin"

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
//	@Router				/databases [get]
func (c *DBController) CreateOrUpdate(ctx *gin.Context) {}
