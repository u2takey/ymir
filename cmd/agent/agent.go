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
			EnvVar: "MASTER_ADDR",
			Name:   "master-addr",
			Usage:  "master addr",
			Value:  "http://ymir-server:5600",
		},
		cli.DurationFlag{
			EnvVar: "DEFAULT_TIMEOUT",
			Name:   "default-timeout",
			Usage:  "default timeout of task",
			Value:  5 * time.Minute,
		},
		cli.StringFlag{
			EnvVar: "JOB_NAME",
			Name:   "job-name",
			Usage:  "job name",
		},
		cli.StringFlag{
			EnvVar: "WORK_ID",
			Name:   "work-id",
			Usage:  "work id",
		},
		cli.StringFlag{
			EnvVar: "INSTANCE_ID",
			Name:   "instance-id",
			Usage:  "instance id",
		},
		cli.StringFlag{
			EnvVar: "NODE_NAME",
			Name:   "node-name",
			Usage:  "node name",
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
	cfg.MasterAddr = c.String("master-addr")
	cfg.TaskSetTimeout = c.Duration("default-timeout")
	cfg.JobName = c.String("job-name")
	cfg.WorkID = c.String("work-id")
	cfg.InstanceID = c.String("instance-id")
	cfg.NodeName = c.String("node-name")
	agent := agent.New(cfg)
	agent.RunTasks()
	return nil
}
