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
