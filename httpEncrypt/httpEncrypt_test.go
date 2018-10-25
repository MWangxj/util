package httpEncrypt

import (
	"fmt"
	"git.dian.so/leto/util/byte2str"
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
	if res, err = Get(app, "192.168.49.97:8080/demo", nil, mm); err != nil {
		t.Fail()
		return
	}
	fmt.Println(byte2str.BytesToString(res))
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
	if res, err = Post(app, "192.168.49.97:8080/postDemo", nil, mm); err != nil {
		t.Fail()
		return
	}
	fmt.Println(byte2str.BytesToString(res))
}

func TestParseUrl(t *testing.T) {
	l := "www.baidu.com/test/api"
	u, err := url.Parse(l)
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.Opaque)
	fmt.Println(u.RawPath)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Scheme)
}
