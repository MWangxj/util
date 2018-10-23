package http

import "testing"

func TestGet(t *testing.T) {
	Get("www.baidu.com",nil,nil,nil)
}

func TestPost(t *testing.T) {
	Post("www.baidu.com",nil,nil,nil)
}
