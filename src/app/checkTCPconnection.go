package main

import (
	"net"
	"time"
)

//端口连接检测，连接成功err为空，字符串拼接1，反之0，将值传入管道
func checkTCP(ip string, port string, timeout int, name string) {
	//defer func() { <-blocker }()
	defer wg.Done()
	var index string
	connection, err := net.DialTimeout("tcp", ip+":"+port, time.Duration(timeout)*time.Millisecond)

	if err == nil {
		index = "check_port{name=" + `"` + name + `"` + "}" + " 1"
		pipeport <- index
		connection.Close()
	} else {
		index = "check_port{name=" + `"` + name + `"` + "}" + " 0"
		pipeport <- index
	}
}
