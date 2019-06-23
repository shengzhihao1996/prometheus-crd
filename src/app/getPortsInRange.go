package main

import (
	"errors"
	"strconv"
)

//获取端口区间，配置文件端口可以100-200这样写。
func getPortsInRange(rangeStr string) ([]int, error) {
	ports := []int{}

	portNum, err := strconv.Atoi(rangeStr)
	if err != nil {
		return ports, errors.New("Port Parse Error")
	}
	if isValidPortNumber(portNum) {
		ports = append(ports, portNum)
	}

	return ports, nil
}

//端口判定，1--65535
func isValidPortNumber(n int) bool {
	if 0 < n && n <= 65535 {
		return true
	}
	return false
}
