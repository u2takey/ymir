package utils

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"
)

func DebugJson(in interface{}) {
	buf, _ := json.MarshalIndent(in, "", "	")
	logrus.Debug(string(buf))
}
