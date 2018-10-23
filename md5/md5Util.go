package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Genmd5String(b []byte) string {
	md := md5.New()
	md.Write(b)
	return hex.EncodeToString(md.Sum(nil))
}
