package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	ip, err := myIp()
	if err != nil {
		panic(err)
	}
	fmt.Println(ip)
}

func myIp() (string, error) {
	c, err := net.Dial("tcp", "ns1.dnspod.net:6666")
	if err != nil {
		return "", err
	}
	defer c.Close()
	c.SetDeadline(time.Now().Add(time.Second * 5))
	buf := make([]byte, 32)
	n, err := c.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}
