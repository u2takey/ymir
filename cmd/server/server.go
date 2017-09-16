package server

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/arlert/ymir/model"
	router "github.com/arlert/ymir/server"
)

// Command exports the server command.
var Command = cli.Command{
	Name:   "server",
	Usage:  "starts the ymir server daemon",
	Action: server,
	Flags: []cli.Flag{
		cli.BoolFlag{
			EnvVar: "DEBUG",
			Name:   "debug",
			Usage:  "start the server in debug mode",
		},
		cli.StringFlag{
			EnvVar: "SERVER_ADDR",
			Name:   "server-addr",
			Usage:  "server address",
			Value:  ":5600",
		},
		cli.StringFlag{
			EnvVar: "KUBERNETE_ADDR",
			Name:   "kubernete-addr",
			Usage:  "kubernete addr",
			Value:  "https://kubernetes.default",
		},
	},
}

func server(c *cli.Context) error {
	// debug level if requested by user
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}
	cfg := &model.ServerConfig{}
	cfg.KubeAddr = c.String("kubernete-addr")
	// setup the server and start the listener
	handler := router.Load(cfg)

	// start the server
	return http.ListenAndServe(
		c.String("server-addr"),
		handler,
	)
}
