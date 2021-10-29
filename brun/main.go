package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	_ "han-xuefeng/gin-scaffold"
	"han-xuefeng/gin-scaffold/infra"
)

func main() {

	file := kvs.GetCurrentFilePath("config.ini", 1)
	// 加载和解析配置
	conf := ini.NewIniFileConfigSource(file)
	app := infra.New(conf)
	app.Start()
}
