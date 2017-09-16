package task

import (
	"fmt"
	"time"
)

// TTaskSet ...
type TTaskSet interface {
	Name() string
	Weight() int // default 1
	Tasks() []TTask
}

//TTask ....
type TTask interface {
	Name() string
	//todo TimeOunt()time.Duration
	Run() int
}

var tasksets = make(map[string]NewFunc, 0)
var runtime = time.Second * 0
var routineCount = 100

// NewFunc ...
type NewFunc func() TTaskSet

// Register ...
func Register(newfunc NewFunc) {
	t := newfunc()
	_, registered := tasksets[t.Name()]
	if registered {
		panic(fmt.Sprintf("TTask named %s already registered", t.Name()))
	}
	tasksets[t.Name()] = newfunc
}

// SetRunTime ...
func SetRunTime(t time.Duration) {
	runtime = t
}

// GetRunTime ...
func GetRunTime() time.Duration {
	return runtime
}

// SetRoutineCount ...
func SetRoutineCount(c int) {
	routineCount = c
}

// GetRoutineCount ...
func GetRoutineCount() int {
	return routineCount
}

// Sets ...
func Sets() map[string]NewFunc {
	return tasksets
}
