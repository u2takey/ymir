package model

import "time"

// TJob is a test job config
type TJob struct {
	Name        string
	Description string
	Tasks       []TTask
	Created     time.Time
	Updated     time.Time
}

// UserSetting is enduser count
type UserSetting struct {
	UserStart     int
	Step          int
	StageDuration int //s
	StageCount    int
	Interval      int //ms
	TimeOut       int //ms
}

// TTask ...
type TTask struct {
	Name     string
	Requests []TRequest
	Weight   int
}

// HostSetting is for test Host
// type HostSetting struct {
// 	IP   string
// 	Host string
// }

// TRequest is test request
type TRequest struct {
	URL      string
	Method   string
	Headers  map[string]string
	Body     []byte
	Variable *VariableGetter
	RetCheck *RetChecker
}

// RetChecker ...
type RetChecker struct {
	Getter VariableGetter
	Expect string
}

// VariableGetter get variable from body or header
type VariableGetter struct {
	Name       string
	FromType   string
	FromatType string
	Getter     string
}
