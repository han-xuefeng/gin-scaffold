package main

import (
	"han-xuefeng/gin-scaffold/infra"
	"han-xuefeng/gin-scaffold/infra/base"
)

func main() {
	infra.Register(&base.GinServerStarter{})
	infra.SystemRun()
}
