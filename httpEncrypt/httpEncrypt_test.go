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
	header := map[string]string{
		"s": app.source,
	}
	var res map[string]interface{}
	if err := Get(app, "192.168.49.97:8080", "/demo", header, mm, &res); err != nil {
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
	header := map[string]string{
		"s": app.source,
	}
	if err := Post(app, "192.168.49.97:8080", "/postDemo", header, mm, &res); err != nil {
		t.Fail()
		return
	}
	fmt.Println(res)
}
