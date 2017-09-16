package utils

import (
	"encoding/json"
	"fmt"
)

// PrintJSON ...
func PrintJSON(in interface{}) {
	buf, _ := json.MarshalIndent(in, "", "	")
	fmt.Println(string(buf))
}
