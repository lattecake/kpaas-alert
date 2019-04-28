package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lattecake/kpaas-alert/pkg/config"
)
import _ "github.com/go-sql-driver/mysql"

var gdb *gorm.DB

func GetMysqlSession() *gorm.DB {

	if gdb != nil {
		return gdb
	}

	gdb, err := gorm.Open(config.Storage.Db.Driver, config.Storage.Db.DbUrl)
	if err != nil {
		fmt.Println(config.Storage.Db.Driver, config.Storage.Db.DbUrl)
		return nil
	}

	defer func() {
		if err = gdb.Close(); err != nil {
			panic(err)
		}
	}()

	gdb.AutoMigrate(&Alert{})

	return gdb
}
