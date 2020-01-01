package net

import (
	"net"
	"strconv"
	"strings"
)

func ToIPv4(ip string) net.IP {
	marks := strings.Split(ip, ".")
	if len(marks) < 4 {
		return net.IPv4zero
	}

	a, _ := strconv.Atoi(marks[0])
	b, _ := strconv.Atoi(marks[1])
	c, _ := strconv.Atoi(marks[2])
	d, _ := strconv.Atoi(marks[3])

	return net.IPv4(byte(a), byte(b), byte(c), byte(d))
}

func IsPrivateIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return true
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return true
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return true
		case ip4[0] == 192 && ip4[1] == 168:
			return true
		default:
			return false
		}
	}
	return false
}
