package net_test

import (
	"fmt"
	"github.com/zhanw15/gopkg/net"
	"testing"
)

func TestIsPrivateIP(t *testing.T) {
	fmt.Println(net.IsPrivateIP(net.ToIPv4("127.0.0.1")))
	fmt.Println(net.IsPrivateIP(net.ToIPv4("128.7.8.8")))
}