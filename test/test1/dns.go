package main

import (
	"net"
	"time"
	"fmt"
	"log"
)

func main() {
	now := time.Now()
	ip, err := net.ResolveIPAddr("ip", "kubernetes.default")
	use := time.Since(now).Nanoseconds()
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("time used: %v, ip: %v\n", use, ip)
}
