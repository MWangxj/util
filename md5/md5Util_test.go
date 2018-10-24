package md5

import (
	"fmt"
	"git.dian.so/leto/util/byte2str"
	"testing"
)

func TestGenmd5String(t *testing.T) {
	s := "guishan@dian.so"
	fmt.Println(Genmd5String(byte2str.StringToBytes(s)))
}

func TestGenmd5String2(t *testing.T) {
	s := `/getbox||0HBA1oqm8/os7XLKP8S7IF4bsxfnQZeCl5XqkHyrCrM=||1539736005110||apoq2rEGljmefWfP||1||apollo`
	fmt.Println(Genmd5String(byte2str.StringToBytes(s)))
}