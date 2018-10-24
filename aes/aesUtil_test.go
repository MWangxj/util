package aes

import (
	"fmt"
	"git.dian.so/leto/util/base64"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	//src := getRepeatString("b", 17)
	//key := getRepeatString("a", 17)
	src := `{"id":"123","name":"张三"}`
	key := "apoq2rEGljmefWfP"
	res, err := AesEncrypt([]byte(src), []byte(key), Aes128)
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(base64.Base64Encoding(res))
}

func getRepeatString(ch string, num int) string {
	res := ""
	for i := 0; i < num; i++ {
		res += ch
	}
	fmt.Println(res)
	return res
}

func TestAesDecrypt(t *testing.T) {
	bstr := "0HBA1oqm8/os7XLKP8S7IMFWAJv/E8BKydJouHWdD8Y="
	key := "apoq2rEGljmefWfP"

	data, err := base64.Base64Decoding(bstr)
	if err != nil {
		t.Fail()
		return
	}
	res, err := AesDecrypt(data, []byte(key), Aes128)
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(string(res))
}
