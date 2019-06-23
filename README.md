## if you want to build image, please install docker, you must have access to an external network, then use: 

```
make install
```

###### 注：程序报错请第一时间查配置文件(必须和例子格式一样)：config。
###### 注：想用什么base镜像，就得在什么base系统构建代码，例如：ubuntu下，代码构建的二进制文件不能再alpine里执行。本次构建基于alpine。
###### 注：修改完配置文件后，该服务可以实时读取配置文件。