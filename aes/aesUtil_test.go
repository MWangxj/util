package aes

import (
	"fmt"
	"git.dian.so/leto/util/base64"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	src := getaesdata("b")
	key := getaesdata("A")
	fmt.Println([]byte(src))
	fmt.Println([]byte(key))
	res, err := AesEncrypt([]byte(src), []byte(key), Aes128)
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(res)
	fmt.Println(base64.Base64Encoding(res))
}

func getaesdata(a string) string {
	res := ""
	for i := 0; i < 16; i++ {
		res += a
	}
	fmt.Println(res)
	return res
}

func TestAesDecrypt(t *testing.T) {
	bstr := "GQ8PMJyhiFU+0Ucv3GSPnwhPjlt67xTNVuhrnTKFiu4="
	key := "zheshimishi"

	data, err := base64.Base64Decoding(bstr)
	if err != nil {
		t.Fail()
		return
	}
	res, err := AesDecrypt(data, []byte(key), Aes256)
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(string(res))
}
