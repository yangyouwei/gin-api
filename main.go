package main

import (
	db "github.com/yangyouwei/gin-api/databases"
	. "github.com/yangyouwei/gin-api/router"
)
func main()  {
	defer db.SqlDB.Close()
	router := InitRouter()
	router.Run(":8080")
}
