package lib

import (
	"fmt"
	"gin_server/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

var dbMysql *xorm.Engine

func mysqlEngine(DbFlag string) *xorm.Engine {
	connectString := helper.Config.GetString("DB.Mysql."+DbFlag+".UserName") +
		":" +
		helper.Config.GetString("DB.Mysql."+DbFlag+".PassWord") +
		"@tcp(" + helper.Config.GetString("DB.Mysql."+DbFlag+".Connection") +
		":" +
		helper.Config.GetString("DB.Mysql."+DbFlag+".Port") + ")/" +
		helper.Config.GetString("DB.Mysql."+DbFlag+".DbName")
	engine, err := xorm.NewEngine(helper.Config.GetString("DB.Mysql."+DbFlag+".DriverName"), connectString)
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
	dbMysql = mysqlEngine("Andoncloud")
	fmt.Println("数据库初始化成功")
}
