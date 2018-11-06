package base64

import (
	"fmt"
	"testing"
)

func TestBase64Encoding(t *testing.T) {
	b := []byte{1,2,3,4,5,6}
	fmt.Println(Base64UrlEncodeing(b))
}

func TestBase64Decoding(t *testing.T) {
	str := "AQIDBAUG"
	fmt.Println(Base64UrlDecoding(str))
}