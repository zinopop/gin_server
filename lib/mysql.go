package lib

import (
	"database/sql"
	"gin_server_v1/helper"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
func init()  {
	connectString := helper.Config.GetString("DB.Mysql.default.UserName")+
		":"+
		helper.Config.GetString("DB.Mysql.default.PassWord")+
		"@tcp("+helper.Config.GetString("DB.Mysql.default.Connection")+
		":"+
		helper.Config.GetString("DB.Mysql.default.Port")+")/"+
		helper.Config.GetString("DB.Mysql.default.DbName")
	db, _ := sql.Open(helper.Config.GetString("DB.Mysql.default.DriverName"), connectString)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	Db = db
}

