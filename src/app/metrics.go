package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

//无限循环，CheckPort获取端口信息传入管道，LenCount获取管道值后初始化清空管道。
func metrics() {
	pipeport = make(chan string, 100)
	pipeurl = make(chan string, 100)
	for {
		var config Config
		start := time.Now().UnixNano()

		yamlFile, _ := filepath.Abs("../../config")
		yamlData, err := ioutil.ReadFile(yamlFile)
		config = Config{}
		err = yaml.Unmarshal(yamlData, &config)

		if err != nil {
			fmt.Printf("error: %v", err)
		}

		itemscount(&config)

		go CheckPort(&config)
		go Checkurl(&config)

		wg.Wait()
		end := time.Now().UnixNano()
		fmt.Printf("coust:%d ms\n", (end-start)/1000000)

		initialize(&config)
		time.Sleep(10 * time.Second)

	}
}
