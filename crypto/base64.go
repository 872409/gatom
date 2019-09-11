package crypto

import "encoding/base64"

func Base64EncodeStr(str string) string {
	return Base64Encode([]byte(str))
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64DecodeStr(str string) (string, error) {
	data, err := Base64Decode(str)
	return string(data), err
}
func Base64Decode(str string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	return data, err
}
