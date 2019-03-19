package httpEncrypt

import (
	"encoding/json"
	"errors"
	"github.com/MWangxj/util/aes"
	"github.com/MWangxj/util/base64"
	"github.com/MWangxj/util/byte2str"
	"github.com/MWangxj/util/http"
	"github.com/MWangxj/util/md5"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type version int

var ver version = 1

var App *app

type app struct {
	source string
	secret string
	salt   string
}

type httpMethod string

const (
	HttpGet  httpMethod = "GET"
	HttpPost            = "POST"
)

func Do(ap *app, method httpMethod, urlStr string, header map[string]string, param interface{}) (resp []byte, err error) {
	if !strings.HasPrefix(urlStr, "http") {
		urlStr = "http://" + urlStr
	}
	switch method {
	case HttpGet:
		return get(ap, urlStr, header, param)
	case HttpPost:
		return post(ap, urlStr, header, param)
	default:
		return nil, errors.New("http method not suport")
	}
}

func NewApp(source, secret, salt string) *app {
	return &app{
		source: source,
		secret: secret,
		salt:   salt,
	}
}

func get(ap *app, urlStr string, header map[string]string, param interface{}) (resp []byte, err error) {
	var (
		newParam map[string]string
		u        *url.URL
	)
	if u, err = url.Parse(urlStr); err != nil {
		return nil, err
	}
	if newParam, err = format(ap, u.Path, param, ver); err != nil {
		return nil, err
	}
	if header == nil {
		header = map[string]string{
			"token": ap.source,
		}
	} else {
		header["token"] = ap.source
	}
	return http.Get(urlStr, header, newParam)
}

func post(ap *app, urlStr string, header map[string]string, param interface{}) (resp []byte, err error) {
	var (
		newParam map[string]string
		u        *url.URL
	)
	if u, err = url.Parse(urlStr); err != nil {
		return nil, err
	}
	if newParam, err = format(ap, u.Path, param, ver); err != nil {
		return nil, err
	}
	if header == nil {
		header = map[string]string{
			"token": ap.source,
		}
	} else {
		header["token"] = ap.source
	}
	return http.Post(urlStr, header, newParam)
}

func format(ap *app, url string, param interface{}, ver version) (map[string]string, error) {
	var (
		jsData  []byte
		ecyData []byte
		md5Str  string
		err     error
	)
	if jsData, err = json.Marshal(param); err != nil {
		return nil, err
	}
	if ecyData, err = aes.AesEncrypt(jsData, byte2str.StringToBytes(ap.secret), aes.Aes128); err != nil {
		return nil, err
	}
	urlByte := byte2str.StringToBytes(url)
	combineByte := byte2str.StringToBytes("||")
	baseByte := byte2str.StringToBytes(base64.Base64UrlEncodeing(ecyData))
	ts := strconv.Itoa(int(time.Now().UnixNano() / 1000))
	tsByte := byte2str.StringToBytes(ts)
	sourceByte := byte2str.StringToBytes(ap.source)
	versionByte := byte2str.StringToBytes(strconv.Itoa(int(ver)))
	saltByte := byte2str.StringToBytes(ap.salt)
	comByte := byte2str.BytesCombine(urlByte, combineByte, baseByte, combineByte, tsByte, combineByte, saltByte, combineByte, versionByte, combineByte, sourceByte)
	md5Str = md5.Genmd5String(comByte)
	newParam := map[string]string{
		"data": base64.Base64UrlEncodeing(ecyData),
		"ts":   ts,
		"sign": md5Str[:16],
		"v":    strconv.Itoa(int(ver)),
	}
	return newParam, nil
}
