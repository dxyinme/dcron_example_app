package controllers

import (
	"github.com/gin-gonic/gin"
)

type ClusterController struct{}

//	cluster list nodes godoc
//
//	@Summary			list nodes
//	@Description  list nodes
//	@Tags					cluster
//	@Accept				json
//	@Produce			json
//	@Success			200 {string} OK
//	@Router				/cluster/nodes [get]
func (c *ClusterController) ListNodes(ctx *gin.Context) {
}
