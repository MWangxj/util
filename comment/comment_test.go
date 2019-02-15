package comment

import (
	"fmt"
	"testing"
)

func TestGetFileNamesByDir(t *testing.T) {
	if fs, err := GetGoFileListByDir("."); err != nil {
		t.Fail()
	} else {
		t.Log(fs)
	}
}

func TestGetFileLine(t *testing.T) {
	GetFileLine("/usr/local/gopath/src/git.dian.so/leto/util/comment/comment.go")
}

func TestParseFuncInfo(t *testing.T)  {
	fi:= parseFuncInfo(`func parseFuncInfo(funcStr string) (fi *funcInfo) {`)
	fmt.Printf("%+v",fi)
}