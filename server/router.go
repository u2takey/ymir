package server

import (
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/arlert/ymir/model"
	"github.com/arlert/ymir/server/middleware/header"
	"github.com/arlert/ymir/server/service"
	_ "github.com/arlert/ymir/utils/loghook"
	"github.com/arlert/ymir/utils/reqlog"
)

// Load loads the router
func Load(cfg *model.ServerConfig) http.Handler {

	logrus.Debugf("\n\nLoad with config:\n %+v\n\n", cfg)

	e := gin.New()
	e.Use(gin.Recovery())

	e.Use(header.NoCache)
	e.Use(header.Secure)
	e.Use(header.Version)
	e.Use(header.Options)
	e.Use(reqlog.ReqLoggerMiddleware(logrus.New(), time.RFC3339, true))

	svc := service.New(cfg)
	e.GET("ping", svc.GetPing)

	ex, _ := os.Executable()
	dir := path.Dir(ex)
	e.LoadHTMLFiles(path.Join(dir, "../index.html"))
	e.Use(historyAPIFallback(), static.Serve("/", service.Frontend("/")))

	v1group := e.Group("/api/v1")
	{
		//-----------------------------------------------------------------
		// job
		v1group.GET("/tjobs", svc.GetJobs)
		v1group.GET("/tjobs/:tjobname", svc.GetJobs)      // get tjob
		v1group.POST("/tjobs", svc.PostJobs)              // add tjob     -> add job (set replica to 0)
		v1group.DELETE("/tjobs/:tjobname", svc.DeleteJob) // delete tjob  -> delete job
		v1group.PUT("/tjobs", svc.PutJob)                 // mod tjob
		v1group.PATCH("/tjobs/:tjobname", svc.PatchJob)   // stop/restart tjob

		//-----------------------------------------------------------------
		// tresult
		v1group.GET("/tresult/:tjobname", svc.GetResult)
		v1group.GET("/tresult/:tjobname/:tworkid", svc.GetResult)         // get tresult(s)
		v1group.POST("/tresult", svc.PostResult)                          // add tresult : agent -> master
		v1group.POST("/tlog/:tjobname/:tworkid/:tinstanceid", svc.GetLog) // get instance log

		//-----------------------------------------------------------------
		v1group.GET("/taskready/:tjobname", svc.GetReady) // ok to start test ?  : agent -> master

		//-----------------------------------------------------------------
		// node
		v1group.GET("/nodes", svc.GetNodes)                        // get nodes
		v1group.GET("/nodesmetrics/:nodename", svc.GetNodeMetrics) // get node monitor
	}

	// e.StaticFile("/favicon.ico", path.Join(dir, "../../dashboard/ymir-ui/dist/favicon.ico"))
	// e.Static("", path.Join(dir, "../../dashboard/ymir-ui/dist"))
	// e.Use(historyAPIFallback())

	return e
}

func historyAPIFallback() gin.HandlerFunc {
	// Serve Statics
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			c.Next()
		}
		var contentType = c.Request.Header.Get("Accept")
		if strings.Contains(contentType, "text/html") {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		}
		c.Next()
	}
}
