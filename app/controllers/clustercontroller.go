package controllers

import (
	"app/internal/crontasks"
	"app/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClusterController struct{}

// cluster get node by node name godoc
//
// @Summary				get current node infomation
// @Description 	get current node infomation
// @Tags					cluster
// @Accept				json
// @Produce				json
// @Success				200 {object} types.Node
// @Router				/cluster/node [get]
func (c *ClusterController) GetNode(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,
		types.Node{
			ID: crontasks.CronTasksContainerUtil{}.I().NodeID(),
		})
}
