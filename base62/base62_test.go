package base62

import (
	"testing"
)

func TestPadding(t *testing.T)  {
	n := 4900
	t.Log(Base62encode(n))
}

func TestBase62Encode(t *testing.T) {

}
