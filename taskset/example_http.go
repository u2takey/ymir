package taskset

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/arlert/ymir/task"
)

func init() {
	task.Register(task.NewFunc(newhttptest))
}

// New ....
func newhttptest() task.TTaskSet {
	return &httptest{}
}

type httptest struct {
}

type httptesttask struct {
}

func (h *httptest) Name() string {
	return "httptest"
}
func (h *httptest) Weight() int {
	return 100
}

func (h *httptest) RunTime() time.Duration {
	return 0
}

func (h *httptest) Tasks() []task.TTask {
	return []task.TTask{&httptesttask{}}
}

func (h *httptesttask) Name() string {
	return "task"
}

func (h *httptesttask) Run() int {
	resp, err := http.Get("http://qq.com")
	if err != nil {
		return 500
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return 600
	}
	return resp.StatusCode
}
