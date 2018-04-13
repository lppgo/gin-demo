package mysql

import (
	"demo/services/hsmysql"
	"github.com/go-xorm/xorm"
)

func getHandle(name string) *xorm.Engine {
	return hsmysql.Conn(name)
}

func GetDefault() *xorm.Engine {
	var dConf = "change"
	return hsmysql.Conn(dConf)
}

func GetCust(name string) *xorm.Engine {
	return hsmysql.Conn(name)
}
