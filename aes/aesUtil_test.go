package aes

import (
	"encoding/json"
	"fmt"
	"github.com/MWangxj/util/base64"
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
	fmt.Println(base64.Base64UrlEncodeing(res))
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
	//bstr := "0HBA1oqm8/os7XLKP8S7IMFWAJv/E8BKydJouHWdD8Y="
	bstr := "ch6qTnjwddfqkNyr6b41PexsqxMzVQuDhvLS+nNqURTUgWPbNepKdwPJorOxg6EfXYPEfS/76xHfEC9TOJvC4tstJk37DdpSUbNKuhfCjFHpAtbAzV+2VLfGizskrpUQMg/NqrHc8VQ0h1kNF/p2dHD5HPg0LhCf2rTi/O+F4JPjpJhepqERROJxtGw0Vhx5tCpco0lDfuSWwAPtfnPaIA=="
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
	mm := make(map[string]interface{})
	if err := json.Unmarshal(res,&mm);err!=nil{
		t.Fail()
		t.Log(err)
		return
	}
	fmt.Println(mm)
}

func TestAesEncrypt2(t *testing.T) {
	str := "YQvFyy5dIshSgxsAXgWg9IjkzoC1WH4iL9weidNDHbLQ2w0aK+I1OD7QcSgb0W3e+Md5+o3snDfk/zyOxFTMOvQeADeqW9yWSc44Vs2HdinhZNR52h6UWL7s6MHsqV9qvl4vXwvismtQFXZQGHVfv31lkdUrDocHB+P7+wnDo8GA4o6y2TI7dycf2TCcdLejavVzPFnQvHQV9grkFwdisbyfwkRsQ9uE4sicHT6NBs8="
	key := "4fe281fac41b9b74"
	data, err := base64.Base64Decoding(str)
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