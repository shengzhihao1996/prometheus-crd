package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

//端口检测，从配置文件获取值，启动checkTCP，检查网络连接。
func CheckPort(c *Config) {

	for _, ipRange := range c.IPRanges {
		rangeStr := ipRange.Range
		nameStr := ipRange.Name
		components := strings.Split(rangeStr, ":")

		if len(components) != 2 {
			fmt.Println("Invalid range string provided:", rangeStr)
		}

		ip, ipnet, err := net.ParseCIDR(components[0])
		if err != nil {
			fmt.Println(err)
		}

		ports, err := getPortsInRange(components[1])
		if err != nil {
			fmt.Println(err)
		}

		for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
			for _, port := range ports {
				go checkTCP(ip.String(), strconv.Itoa(port), c.TimeOut, nameStr)

			}
		}
	}

}

//没看懂，无论如何那个if判断都是对的
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

//从config文件读取url信息，
func Checkurl(c *Config) {
	for _, urlRange := range c.URLRanges {
		nameStr := urlRange.Name
		go exec_shell(nameStr)
	}
}
