package main

import (
	"fmt"
	"os/exec"
)

//执行shell的curl命令取出状态吗
func exec_shell(url string) {
	defer wg.Done()
	var index string
	s := "curl --connect-timeout 3 -skI " + url + "|awk 'NR==1{print $2}'"
	cmd := exec.Command("/bin/sh", "-c", s)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Check your url name, it's not true -> ", s)
	}
	if len(output) != 0 {
		index = "check_url{name=" + `"` + url + `"` + "} " + string(output)
		pipeurl <- index

	} else {
		fmt.Println("error!!!  Check your url name, it's not true -> ", url)
		index = "check_url{name=" + `"` + url + `"` + "} " + "0"
		pipeurl <- index
	}
}
