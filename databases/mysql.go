package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB
//初始化方法
func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:gaopeng@tcp(192.168.2.250:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	//连接检测
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
