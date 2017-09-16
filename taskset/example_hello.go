package taskset

import (
	"time"

	"github.com/arlert/ymir/task"
)

func init() {
	task.Register(task.NewFunc(newhello))
}

// New ....
func newhello() task.TTaskSet {
	return &hello{}
}

type hello struct {
}

type hellotask struct {
}

func (h *hello) Name() string {
	return "hello"
}
func (h *hello) Weight() int {
	return 50
}

func (h *hello) RunTime() time.Duration {
	return 0
}

func (h *hello) Tasks() []task.TTask {
	return []task.TTask{&hellotask{}}
}

func (h *hellotask) Name() string {
	return "task"
}

func (h *hellotask) Run() int {
	time.Sleep(time.Microsecond * 10)
	return 0
}
