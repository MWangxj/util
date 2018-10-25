package http

import (
	"git.dian.so/leto/util/byte2str"
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := Get("www.baidu.com", nil, nil)
	if err != nil {
		t.Fail()
		return
	}
	t.Log(byte2str.BytesToString(resp))
}

func TestPost(t *testing.T) {
	resp, err := Post("www.baidu.com", nil, nil)
	if err != nil {
		t.Fail()
		return
	}
	t.Log(byte2str.BytesToString(resp))
}
