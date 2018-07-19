package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:HX141656@/hx_shiyan?charset=utf8")
	if err != nil {

		fmt.Println("err:", err)

	}

	//关闭数据库，db会被多个goroutine共享，可以不调用
	defer db.Close()

	//查询数据，指定字段名，返回sql.Rows结果集
	rows, err1 := db.Query("select id,name from employee")
	if err1 != nil {

		fmt.Println("err1:", err1)

	}
	id := 0
	name := ""
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}

}
