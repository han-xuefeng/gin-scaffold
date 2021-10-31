package gin_scaffold

import (
	"han-xuefeng/gin-scaffold/infra"
	"han-xuefeng/gin-scaffold/infra/base"
)

//这里是注册所有starter

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
}
