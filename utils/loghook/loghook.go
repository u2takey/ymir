package loghook

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

func init() {
	logrus.StandardLogger().Hooks.Add(ContextHook{})
}

type ContextHook struct{}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook ContextHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 4, 4)
	cnt := runtime.Callers(6, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !strings.Contains(name, "github.com/Sirupsen/logrus") {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data["where"] = fmt.Sprintf("%s:%d", filepath.Base(file), line)
			break
		}
	}
	return nil
}
