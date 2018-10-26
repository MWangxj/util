package commonEncrypt

import (
	"encoding/json"
	"git.dian.so/leto/util/aes"
	"git.dian.so/leto/util/byte2str"
	"git.dian.so/leto/util/md5"
	"strings"
)

func Encrypt(obj interface{}, accessKey string) (encrypt []byte, err error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	b, err := aes.AesEncrypt(data, byte2str.StringToBytes(accessKey), aes.Aes128)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Sign(url, encryptData, timeStamp, salt, v, source string) string {

	data := []string{url, encryptData, timeStamp, salt, v, source}
	dataStr := strings.Join(data, "||")
	dataByte := byte2str.StringToBytes(dataStr)
	return md5.Genmd5String(dataByte)
}

func Decrypt(encryptData []byte, aesKey string) ([]byte, error) {
	return aes.AesDecrypt(encryptData, byte2str.StringToBytes(aesKey), aes.Aes128)
}

func VerifySign(url, sign, ts, encryptData, v, token, salt string) bool {
	signStr := Sign(url, encryptData, ts, salt, v, token)
	if signStr != sign {
		return false
	}
	return true
}
