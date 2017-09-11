package utils

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	defaultQPS   = 1e6
	defaultBurst = 1e6
)

func CreateK8sClientByConfig(cfg *rest.Config) (*kubernetes.Clientset, error) {
	if cfg.QPS == 0 {
		cfg.QPS = defaultQPS
	}
	if cfg.Burst == 0 {
		cfg.Burst = defaultBurst
	}
	if cfg.ContentType == "" {
		cfg.ContentType = "application/vnd.kubernetes.protobuf"
	}
	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}
