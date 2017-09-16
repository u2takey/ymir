package agent

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/arlert/ymir/agent"
	"github.com/arlert/ymir/model"
)

// Command exports the server command.
var Command = cli.Command{
	Name:   "agent",
	Usage:  "starts the ymir agent daemon",
	Action: server,
	Flags: []cli.Flag{
		cli.BoolFlag{
			EnvVar: "DEBUG",
			Name:   "debug",
			Usage:  "start the agent in debug mode",
		},
		cli.StringFlag{
			EnvVar: "KUBERNETE_ADDR",
			Name:   "kubernete-addr",
			Usage:  "kubernete addr",
			Value:  "https://kubernetes.default",
		},
		cli.DurationFlag{
			EnvVar: "DEFAULT_TIMEOUT",
			Name:   "default-timeout",
			Usage:  "default timeout of task",
			Value:  5 * time.Minute,
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
	cfg := &model.AgentConfig{}
	cfg.KubeAddr = c.String("kubernete-addr")
	cfg.TaskSetTimeout = c.Duration("default-timeout")
	agent := agent.New(cfg)
	agent.RunTasks()
	return nil
}
