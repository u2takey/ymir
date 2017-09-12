package server

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"github.com/arlert/ymir/model"
	// "github.com/arlert/malcolm/server/service"
	"github.com/arlert/ymir/server/middleware/header"
	_ "github.com/arlert/ymir/utils/loghook"
	"github.com/arlert/ymir/utils/reqlog"
)

// Load loads the router
func Load(cfg *model.Config) http.Handler {

	logrus.Debugf("\n\nLoad with config:\n %+v\n\n", cfg)

	e := gin.New()
	e.Use(gin.Recovery())

	e.Use(header.NoCache)
	e.Use(header.Secure)
	e.Use(header.Version)
	e.Use(header.Options)
	e.Use(reqlog.ReqLoggerMiddleware(logrus.New(), time.RFC3339, true))

	// svc := service.New(cfg)
	//svc := &service.Service{}

	e.GET("ping", svc.GetPing)

	//e.Use(static.Serve("/", utils.Frontend("dist")))

	return e
}
