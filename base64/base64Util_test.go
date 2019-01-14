package base64

import (
	"fmt"
	`github.com/GiterLab/gomathbits`
	"testing"
)

func TestBase64Encoding(t *testing.T) {
	b := []byte{1, 2, 3, 4, 5, 6}
	fmt.Println(Base64UrlEncodeing(b))
}

func TestBase64UrlEncodeing(t *testing.T) {
	a := gomathbits.Uint32ToBytes(uint32(1))
	s :=Base64UrlEncodeing(a)
	t.Log(s)
	b ,_:=Base64UrlDecoding(s)
	u,_ := gomathbits.BytesToUint32(b)
	t.Log(u)
}

func TestBase64Decoding(t *testing.T) {
	str := "AQIDBAUG"
	fmt.Println(Base64UrlDecoding(str))
}
