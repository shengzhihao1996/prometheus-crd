install: go-build docker-build
go-build:
	mkdir bin
	docker run  --rm -v `pwd`:/go registry.cn-huhehaote.aliyuncs.com/shengzhihao/go-build:v1  go get -v gopkg.in/yaml.v2
	docker run  --rm -v `pwd`:/go registry.cn-huhehaote.aliyuncs.com/shengzhihao/go-build:v1 go build -o /go/bin/app /go/src/app
docker-build:
	docker build -t prometheus-addon:release-1.0.0 .
