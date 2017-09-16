package model

import "time"

// ServerConfig is server config
type ServerConfig struct {
	KubeAddr string
}

// AgentConfig ...
type AgentConfig struct {
	KubeAddr       string
	TaskSetTimeout time.Duration
}
