package base

import (
	"github.com/tietang/props/kvs"
	"han-xuefeng/gin-scaffold/infra"
)

// props ćšć±ćé
var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter)Init(ctx infra.StarterContext) {
	props = ctx.Props()
}
