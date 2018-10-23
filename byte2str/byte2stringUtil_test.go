package byte2str

import (
	"fmt"
	"testing"
)

func TestByteDelZero(t *testing.T) {
	b := []byte{1, 2, 3, 40, 0, 0, 0, 0}
	fmt.Println(ByteDelZero(b))
}

func TestBytesToString(t *testing.T) {
	b := StringToBytes("wangxianjin")
	fmt.Println(BytesToString(b))
}
