package util

import "encoding/json"

func ToJsonString(object interface{}) string {
	js, _ := json.Marshal(object)
	return string(js)
}
func FromJson(jsonStr string, object interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), object)
	return err
}
