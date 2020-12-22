package model

import (
	"database/sql"
	"gin_server_v1/lib"
)

func GetById(id int) *sql.Rows {
	rows,_ := lib.Db.Query("select * from zhaopin_user")
	rows.Close()
	return rows
}
