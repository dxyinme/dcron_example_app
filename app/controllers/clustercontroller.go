package controllers

import (
	"app/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClusterController struct{}

// cluster list nodes godoc
//
// @Summary				list nodes
// @Description 	list nodes
// @Tags					cluster
// @Accept				json
// @Produce				json
// @Success				200 {object} []types.Node
// @Router				/cluster/nodes [get]
func (c *ClusterController) ListNodes(ctx *gin.Context) {
	ListNodesResp := make([]types.Node, 0)

	ctx.JSON(http.StatusOK, ListNodesResp)
}

// cluster get node by node name godoc
//
// @Summary				get node
// @Description 	get node
// @Tags					cluster
// @Accept				json
// @Produce				json
// @Success				200 {object} types.Node
// @Router				/cluster/nodes/{nodeName} [get]
func (c *ClusterController) GetNode(ctx *gin.Context) {
}
