package byte2str

import (
	"bytes"
	"unsafe"
)

// BytesCombine 合并byte数组
func BytesCombine(pBytes ...[]byte) []byte {
	lenth := len(pBytes)
	s := make([][]byte, lenth)
	for i, v := range pBytes {
		s[i] = v
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

// StringToBytes string类型转化为[]byte
func StringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// BytesToString []byte转化为string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// ByteDelZero 截取到[]byte中的结束符
func ByteDelZero(b []byte) []byte {
	for i, v := range b {
		if v == 0 {
			return b[:i]
		}
	}
	return b
}
