package util

import "net"

type List interface {
	Append(val interface{}) error
	Length() int
	Index(val interface{}) int
}

type StringArray struct {
}

func Index(arr []interface{}, val interface{}) int {
	for i, ele := range arr {
		if ele == val {
			return i
		}
	}

	return -1
}

func GetIPAddr() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, ip := range addrs {
			if ipnet, ok := ip.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String(), nil
				}
			}
		}
	}

	return "", nil
}
