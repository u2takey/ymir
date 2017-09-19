package service

import (
	"github.com/gin-gonic/gin"

	req "github.com/arlert/ymir/utils/reqlog"
)

// GetReady : ok to start test ?  : agent -> master
func (s *Service) GetReady(c *gin.Context) {
	req.Entry(c).Debug("GetReady")
	c.String(200, "to be implemented")
}
