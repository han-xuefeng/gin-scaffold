package base

import (
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
	"han-xuefeng/gin-scaffold/infra"
)

// dbx 全局暴露的数据库实例 先用这个dbx  后续在改

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
	log.Info("数据库连接成功")
	database = dbx
}

