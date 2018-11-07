package base62
/*
import (
	"github.com/GiterLab/gomathbits"
	"math"
)

var base = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

//func padding(b *[]byte) {
//	if len(*b)%4 != 0 {
//		 4-len(*b)%4)
//		*b = append(*b, ...)
//	}
//	return
//}

func Base62Encode(b []byte) string {
	padding(&b)
	res := ""
	for i := 0; i < len(b)/4; i += 4 {
		u32, _ := gomathbits.BytesToUint32(b[i : i+4])
		res += base62encode(int(u32))
	}
	return res
}

func base62encode(num int) string {
	baseStr := ""
	for {
		if num <= 0 {
			break
		}

		i := num % 62
		baseStr += base[i]
		num = (num - i) / 62
	}
	return baseStr
}

func Base62Decode(str string) []byte {
	for _, s := range str {
		base62encode(string(s))
	}
}

func base62decode(base62 string) int {
	rs := 0
	len := len(base62)
	f := flip(base)
	for i := 0; i < len; i++ {
		rs += f[string(base62[i])] * int(math.Pow(62, float64(i)))
	}
	return rs
}

func flip(s []string) map[string]int {
	f := make(map[string]int)
	for index, value := range s {
		f[value] = index
	}
	return f
}
*/