package base64

import (
	"encoding/base64"
	"errors"
)

// Base64Encoding []byte序列化成base64string
func Base64Encoding(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}

// Base64Decoding string 反序列化成[]byte
func Base64Decoding(s string) ([]byte, error) {
	if len(s) == 0 {
		return nil, errors.New("src must be not null")
	}
	return base64.StdEncoding.DecodeString(s)
}
