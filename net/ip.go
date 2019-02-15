package net

import (
	`fmt`
	`net`
	`os`
)

func GetIntranetIp() net.IP {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		os.Exit(1)
	}

	for _, address := range addrs {

		fmt.Println(address.String())
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP
			}

		}
	}
	return nil
}

func GetIntranetIpV4() net.IP {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.To4()
			}

		}
	}
	return nil
}
func GetIntranetIpV16() net.IP {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To16() != nil {
				return ipnet.IP.To16()
			}

		}
	}
	return nil
}
