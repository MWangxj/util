## go util 使用

go get -u -v git.dian.so/leto/util


## aes 加密

    import (
        git.dian.so/leto/util/aes
    )

    aes.AesEncrypt(data, key []byte, t AesType) (res []byte, err error)

    aes.AesDecrypt(data, key []byte, t AesType) (res []byte, err error)



## base64 编码

    import (
         git.dian.so/leto/util/base64
    )

    base64.Base64Encoding(b []byte) string

    base64.Base64Decoding(s string) ([]byte, error)


## byte2str byte数组和string相互转化

    import (
         git.dian.so/leto/util/byte2str
    )

    byte2str.BytesCombine(pBytes ...[]byte) []byte

    byte2str.StringToBytes(s string) []byte

    byte2str.BytesToString(b []byte) string

    byte2str.ByteDelZero(b []byte) []byte


## cfg 读取配置包

    import(
        git.dian.so/leto/util/cfg
    )

    cfg.Load(filePath string, cfg interface{}) error

## httpEncrypt http加密请求

    1、go get -u -v git.dian.so/leto/util

    2、import(
        "git.dian.so/leto/util/httpEncrypt"
    )

    3、使用

        httpEncrypt.NewApp (source, secret, salt string) *app

        httpEncrypt.Do (ap *app, method HttpMethod ,urlStr string, param interface{}, ver version) (resp []byte,err error)

## commonEncrypt 手动加密数据


    import(
        git.dian.so/leto/util/commonEncrypt
    )

    commonEncrypt.Encrypt(obj interface{}, accessKey string) (encrypt string, err error)
    commonEncrypt.Sign(url, encryptData, timeStamp, salt, v, source string) string
    commonEncrypt.Decrypt(encryptData, aesKey string) (map[string]interface{}, error)
    commonEncrypt.VerifySign(url, sign, ts, encryptData, v, token, salt string) bool