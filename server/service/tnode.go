package service

import (
	"github.com/gin-gonic/gin"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	req "github.com/arlert/ymir/utils/reqlog"
)

// GetNodes : get nodes
func (s *Service) GetNodes(c *gin.Context) {
	req.Entry(c).Debug("GetNodes")
	nodelsit, err := s.engine.Core().Nodes().List(meta_v1.ListOptions{})
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(200, nodelsit)
}

// GetNodeMetrics : get node monitor
func (s *Service) GetNodeMetrics(c *gin.Context) {
	req.Entry(c).Debug("GetNodeMetrics")
	c.String(200, "to be implemented")
}
