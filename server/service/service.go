package service

import (
	"io/ioutil"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/arlert/ymir/model"
	"github.com/arlert/ymir/utils/k8s"
	req "github.com/arlert/ymir/utils/reqlog"
)

var (
	bearer_token_file = "/var/run/secrets/kubernetes.io/serviceaccount/token"
)

// Service ...
type Service struct {
	config *model.ServerConfig
	engine *kubernetes.Clientset
}

// New ...
func New(cfg *model.ServerConfig) *Service {
	token := ""
	if bearer_token_file != "" {
		bf, err := ioutil.ReadFile(bearer_token_file)
		if err != nil {
			logrus.Error("read bearer_token err ", err)
		}
		token = string(bf)
	}
	if !strings.HasPrefix(cfg.KubeAddr, "http") {
		cfg.KubeAddr = "http://" + cfg.KubeAddr
	}
	resconfig := &rest.Config{
		Host:        cfg.KubeAddr,
		BearerToken: token,
	}
	resconfig.Insecure = true
	client, err := k8s.CreateK8sClientByConfig(resconfig)
	if err != nil {
		logrus.Fatalln("CreateK8sClientByConfig fail", err)
	}
	svc := &Service{
		config: cfg,
		engine: client,
	}
	return svc
}

// GetPing ...
func (s *Service) GetPing(c *gin.Context) {
	req.Entry(c).Debug("GetPing")
	c.String(200, "pong")
}
