package main

import (
	"iip/conf"
	"iip/routes"
)

func main() {

	//从配置文件读入配置
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)

}
