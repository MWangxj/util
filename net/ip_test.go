package net

import (
	`strings`
	`testing`
)

func TestGetIntranetIp(t *testing.T) {
	ip := GetIntranetIp()
	t.Log(ip.String())
	ip = ip.To4()
	t.Log(string([]byte(ip)))
}

func TestStrings(t *testing.T)  {
	s := " Hello !!!!!"
	ts := strings.TrimRight(s, "!")
	t.Log(ts)
}
