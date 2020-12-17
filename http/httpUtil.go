package http

import (
	"encoding/json"
	"github.com/MWangxj/util/byte2str"
	"io/ioutil"
	"net/http"
	"strings"
        "time"
)

// Get
func Get(url string, header, param map[string]string) (resp []byte, err error) {
	var (
		client = &http.Client{}
		req    = &http.Request{}
		res    = &http.Response{}
		protol = "http://"
	)
	pstr := "?"
	for k, v := range param {
		pstr += k + "=" + v + "&"
	}
	pstr = pstr[:len(pstr)-1]
	if !strings.HasPrefix(url, "http") {
		url = protol + url + pstr
	} else {
		url += pstr
	}
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	if res, err = client.Do(req); err != nil {
		return nil, err
	}
	if resp, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
	return resp, nil
}

// Post
func Post(url string, header map[string]string, payload interface{}) (resp []byte, err error) {
	var (
		client = &http.Client{Timeout: 5 * time.Second}
		req    = &http.Request{}
		res    = &http.Response{}
		data   = make([]byte, 0)
		protol = "http://"
	)
	if !strings.HasPrefix(url, "http") {
		url = protol + url
	}
	if b,ok:=payload.([]byte);ok{
		data=b
	}else if data, err = json.Marshal(payload); err != nil {
		return nil, err
	}
	if req, err = http.NewRequest("POST", url, strings.NewReader(byte2str.BytesToString(data))); err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header.Add(k, v)
	}
	if res, err = client.Do(req); err != nil {
		return nil, err
	}
	if resp, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
	return resp, nil
}
