package base

import (
	"fmt"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"han-xuefeng/gin-scaffold/infra"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter)Init(ctx infra.StarterContext) {
	props = ini.NewIniFileConfigSource("config.ini")
	fmt.Println("初始化配置.")
}
