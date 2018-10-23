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

