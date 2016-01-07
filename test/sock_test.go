package test

import (
	"testing"
	"fmt"
	"net"
)

func TestAddress(t *testing.T) {
	name := "192.168.1.1"
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	} 
}

func TestResolveAddress(t *testing.T) {
	name := "www.baidu.com"
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("The address is ", addr.String())	
}