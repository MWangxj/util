package httpEncrypt

import (
	"encoding/json"
	"git.dian.so/leto/util/aes"
	"git.dian.so/leto/util/base64"
	"git.dian.so/leto/util/byte2str"
	"git.dian.so/leto/util/http"
	"git.dian.so/leto/util/md5"
	"strconv"
	"strings"
	"time"
)

type version int

var ver version = 1

type app struct {
	source string
	secret string
	salt   string
}

func NewApp(source, secret, salt string) *app {
	return &app{
		source: source,
		secret: secret,
		salt:   salt,
	}
}

func Get(ap *app, endpoint, url string, param interface{}, resp interface{}) error {
	var (
		newParam map[string]string
		err      error
	)
	if newParam, err = format(ap, endpoint, url, param, ver); err != nil {
		return err
	}
	if strings.HasPrefix(url, "/") {
		url = endpoint + url
	} else {
		url = endpoint + "/" + url
	}
	return http.Get(url, nil, newParam, resp)
}

func Post(ap *app, endpoint, url string, param interface{}, resp interface{}) error {
	var (
		newParam map[string]string
		err      error
	)
	if newParam, err = format(ap, endpoint, url, param, ver); err != nil {
		return err
	}
	if strings.HasPrefix(url, "/") {
		url = endpoint + url
	} else {
		url = endpoint + "/" + url
	}
	return http.Post(url, nil, newParam, resp)
}

func format(ap *app, endpoint, url string, param interface{}, ver version) (map[string]string, error) {
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
	baseByte := byte2str.StringToBytes(base64.Base64Encoding(ecyData))
	ts := strconv.Itoa(int(time.Now().UnixNano() / 1000))
	tsByte := byte2str.StringToBytes(ts)
	sourceByte := byte2str.StringToBytes(ap.source)
	versionByte := byte2str.StringToBytes(strconv.Itoa(int(ver)))
	saltByte := byte2str.StringToBytes(ap.salt)
	comByte := byte2str.BytesCombine(urlByte, combineByte, baseByte, combineByte, tsByte, combineByte, saltByte, combineByte, versionByte, combineByte, sourceByte)
	md5Str = md5.Genmd5String(comByte)
	newParam := map[string]string{
		"data": base64.Base64Encoding(ecyData),
		"ts":   ts,
		"sign": md5Str,
		"v":    strconv.Itoa(int(ver)),
	}
	return newParam, nil
}
