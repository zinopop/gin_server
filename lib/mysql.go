package lib

import (
	"fmt"
	"gin_server_v1/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

var dbMysql *xorm.Engine

func mysqlEngine() *xorm.Engine {
	connectString := helper.Config.GetString("DB.Mysql.default.UserName") +
		":" +
		helper.Config.GetString("DB.Mysql.default.PassWord") +
		"@tcp(" + helper.Config.GetString("DB.Mysql.default.Connection") +
		":" +
		helper.Config.GetString("DB.Mysql.default.Port") + ")/" +
		helper.Config.GetString("DB.Mysql.default.DbName")
	engine, err := xorm.NewEngine(helper.Config.GetString("DB.Mysql.default.DriverName"), connectString)
	if err != nil {
		fmt.Println("mysql初始化失败")
		os.Exit(1)
	}
	//日志打印SQL
	engine.ShowSQL(true)
	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	engine.SetMaxOpenConns(5)
	return engine
}

func GetConn() *xorm.Engine {
	return dbMysql
}

func Close() {
	dbMysql.Close()
}

func init() {
	dbMysql = mysqlEngine()
}
