package conf

import (
	"bytes"
	"encoding/json"
)

// PrettyJSON returns indented JSON data.
func PrettyJSON(data string) string {
	var pretty bytes.Buffer
	err := json.Indent(&pretty, []byte(data), "", "\t")
	if err != nil {
		return data
	}
	return pretty.String()
}
