package http

import (
	"encoding/json"
	"errors"
	"git.dian.so/leto/util/byte2str"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

// Get
func Get(url string, header, param map[string]string, resp interface{}) error {
	var (
		client = &http.Client{}
		req    = &http.Request{}
		res    = &http.Response{}
		data   = make([]byte, 0)
		protol = "http://"
		err    error
	)
	if reflect.ValueOf(resp).Kind() != reflect.Ptr {
		return errors.New("resp interface must be a pointer")
	}
	pstr := "?"
	for k, v := range param {
		pstr += k + "=" + v + "&"
	}
	pstr = pstr[:len(pstr)-1]
	if !strings.HasPrefix(url, "http") {
		url = protol + url + pstr
	}
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	if res, err = client.Do(req); err != nil {
		return err
	}
	if data, err = ioutil.ReadAll(res.Body); err != nil {
		return err
	}
	if err = json.Unmarshal(data, resp); err != nil {
		return err
	}
	return nil
}

// Post
func Post(url string, header map[string]string, payload interface{}, resp interface{}) error {
	var (
		client = &http.Client{}
		req    = &http.Request{}
		res    = &http.Response{}
		data   = make([]byte, 0)
		protol = "http://"
		err    error
	)
	if reflect.ValueOf(resp).Kind() != reflect.Ptr {
		return errors.New("resp interface must be a pointer")
	}
	if !strings.HasPrefix(url, "http") {
		url = protol + url
	}
	if data, err = json.Marshal(payload); err != nil {
		return err
	}
	if req, err = http.NewRequest("POST", url, strings.NewReader(byte2str.BytesToString(data))); err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	if res, err = client.Do(req); err != nil {
		return err
	}
	if data, err = ioutil.ReadAll(res.Body); err != nil {
		return err
	}
	if err = json.Unmarshal(data, resp); err != nil {
		return err
	}
	return nil
}
