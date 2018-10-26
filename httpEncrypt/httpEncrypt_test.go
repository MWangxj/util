package httpEncrypt

import (
	"encoding/json"
	"fmt"
	"git.dian.so/leto/util/base64"
	"git.dian.so/leto/util/byte2str"
	"git.dian.so/leto/util/commonEncrypt"
	"net/http"
	"net/url"
	"testing"
)

func TestGet(t *testing.T) {
	app := NewApp("apollo", "apoq2rEGljmefWfP", "apoq2rEGljmesalt")
	mm := map[string]string{
		"key":  "test",
		"name": "guishan",
	}
	var (
		res []byte
		err error
	)
	if res, err = Do(app, HttpGet, "0.0.0.0:8080/test", nil, mm); err != nil {
		t.Fail()
		return
	}
	fmt.Println(byte2str.BytesToString(res))
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		app := NewApp("apollo", "apoq2rEGljmefWfP", "apoq2rEGljmesalt")
		mm := map[string]string{
			"key":  "test",
			"name": "guishan",
		}
		Do(app, HttpPost, "192.168.49.97:8080/demo", nil, mm)
	}
}

func BenchmarkPost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		app := NewApp("apollo", "apoq2rEGljmefWfP", "apoq2rEGljmesalt")
		mm := map[string]string{
			"key":  "test",
			"name": "guishan",
		}
		Do(app, HttpPost, "127.0.0.1:8080/test", nil, mm)
	}
}

func TestPost(t *testing.T) {

	app := NewApp("apollo", "apoq2rEGljmefWfP", "apoq2rEGljmesalt")
	mm := map[string]string{
		"key":  "test",
		"name": "guishan",
	}
	var (
		res []byte
		err error
	)
	if res, err = Do(app, HttpPost, "127.0.0.1:8080/test", nil, mm); err != nil {
		t.Fail()
		return
	}
	fmt.Println(byte2str.BytesToString(res))
}

func listen() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		b := make([]byte, 1<<11)
		n, _ := request.Body.Read(b)
		b = b[:n]
		mm := make(map[string]string)
		json.Unmarshal(b, &mm)
		data, _ := base64.Base64Decoding(mm["data"])
		t := commonEncrypt.VerifySign(request.URL.Path,mm["sign"],mm["ts"],mm["data"],mm["v"],request.Header.Get("s"),"apoq2rEGljmesalt")
		fmt.Println(t)
		m, err := commonEncrypt.Decrypt(data, "apoq2rEGljmefWfP")
		if err != nil {
			writer.Write(byte2str.StringToBytes(err.Error()))
			return
		}
		writer.Write(m)
	})
	http.ListenAndServe(":8080", nil)
}

func TestListen(t *testing.T) {
	listen()
}

func TestUrlParse(t *testing.T) {
	u := "127.0.0.1:8080/test"
	urlR, err := url.Parse(u)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(fmt.Sprintf("%V", urlR))
}
