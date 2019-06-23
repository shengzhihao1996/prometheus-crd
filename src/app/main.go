package main

import "sync"

//会读取一个配置文件，yaml格式的。
type IPRange struct {
	Range string `yaml:"range"`
	Name  string `yaml:"name"`
}

type URLRange struct {
	Name string `yaml:"name"`
}

type Config struct {
	Concurrency int
	TimeOut     int        `yaml:"timeout"`
	IPRanges    []IPRange  `yaml:"ip"`
	URLRanges   []URLRange `yaml:"url"`
}

var (
	wg sync.WaitGroup
	//map,webui取值输出
	outputport map[string]string
	outputurl  map[string]string
	//channel,为map传值
	pipeport chan string
	pipeurl  chan string
)

//主函数，做webserver。
func main() {
	go metrics()
	webui()
}
