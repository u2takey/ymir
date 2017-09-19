package model

import "time"

var (
	// AppName ...
	AppName = "ymir-app"
	// NodeSelectedKey ...
	NodeSelectedKey = "node-select"
	// DescriptionKey ...
	DescriptionKey = "description"
	// ScriptKey
	ScriptKey = "script"
	WorkKey   = "work"

	TypeTScript = "tscript"
	TypeTResult = "result"
)

// ServerConfig is server config
type ServerConfig struct {
	KubeAddr        string
	TimeoutForStart time.Duration
	JobNamespace    string
	AgentImageName  string
}

// AgentConfig ...
type AgentConfig struct {
	MasterAddr      string
	TaskSetTimeout  time.Duration
	TimeoutForStart time.Duration
	JobName         string
	WorkID          string
	InstanceID      string
	NodeName        string
}
