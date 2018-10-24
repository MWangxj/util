package httpEncrypt

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	app := NewApp("apollo", "apoq2rEGljmefWfP", "apoq2rEGljmesalt")
	mm := map[string]string{
		"key":  "test",
		"name": "guishan",
	}
	var res map[string]interface{}
	if err := Get(app, "192.168.49.97:8080", "/demo", nil, mm, &res); err != nil {
		t.Fail()
		return
	}
	fmt.Println(res)
}

func TestPost(t *testing.T) {

	app := NewApp("apollo", "apoq2rEGljmefWfP", "apoq2rEGljmesalt")
	mm := map[string]string{
		"key":  "test",
		"name": "guishan",
	}
	var res map[string]interface{}
	if err := Post(app, "192.168.49.97:8080", "/postDemo", nil, mm, &res); err != nil {
		t.Fail()
		return
	}
	fmt.Println(res)
}
