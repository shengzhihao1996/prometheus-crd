package main

//初始化map，端口扫描的数据会先进入管道，在这里赋值给map，map在得到值之前会被初始化清空。
func initialize(c *Config) {
	countport := len(c.IPRanges)
	outputport = make(map[string]string)
	for i := 0; i < countport; i++ {
		a := <-pipeport
		outputport[a] = ""
	}
	counturl := len(c.URLRanges)
	outputurl = make(map[string]string)
	for i := 0; i < counturl; i++ {
		a := <-pipeurl
		outputurl[a] = ""
	}

}
