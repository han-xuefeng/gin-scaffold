package base

import (
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
	"han-xuefeng/gin-scaffold/infra"
	_ "github.com/go-sql-driver/mysql"
)

// dbx 全局暴露的数据库实例

var database *dbx.Database

func DbxDatabase() *dbx.Database {
	return database
}

type DbxDatabaseStarter struct {
	infra.BaseStarter
}

func (s DbxDatabaseStarter)Setup(ctx infra.StarterContext){
	conf := Props()
	settings := dbx.Settings{}

	err := kvs.Unmarshal(conf, &settings, "mysql")
	if err != nil {
		panic(err)
	}
	dbx, err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	database = dbx
}

