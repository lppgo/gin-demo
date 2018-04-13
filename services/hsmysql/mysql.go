package hsmysql

import (
	"demo/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var db map[string]*xorm.Engine

// Conn is
func Conn(name string) *xorm.Engine {
	if db == nil {
		db = make(map[string]*xorm.Engine, 2)
	}
	if _, ok := db[name]; ok {
		return db[name]
	}
	dsn, _ := config.GetENVConfigs().String("mysql." + name)
	var err error
	db[name], err = xorm.NewEngine("mysql", dsn)
	db[name].SetMaxIdleConns(5000) //TODO::配置化
	db[name].SetMaxOpenConns(5000) //TODO::配置化
	if err != nil {
		db[name].Close()
	}
	db[name].ShowSQL(true)
	return db[name]
}

// GetHandle is
func GetHandle(name string) *xorm.Engine {
	return db[name]
}
