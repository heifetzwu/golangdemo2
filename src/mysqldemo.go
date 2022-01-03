package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
use mysql;
create table t1 (
c1 varchar(10));
*/
func mysqldemo() {

	//完整的資料格式連線如下
	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// db, err := sql.Open("mysql", "user:password@/dbname")
	// db, err := sql.Open("mysql", "root@/godb")
	db, err := sql.Open("mysql", "root:password@/mysql")

	if err != nil {
		fmt.Println("err=", err)
		return
	}

	result, err := db.Exec(
		"INSERT INTO t1 (c1) VALUES ( ?)",
		"test",
	)
	fmt.Println("result=", result)

	var c1 string

	rows, err := db.Query("SELECT c1 FROM t1 WHERE c1 = ?", "test")

	if err != nil {
		fmt.Println("err=", err)
		return
	}

	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&c1); err != nil {
			fmt.Println("err =", err)
			return
		}
		fmt.Printf("c1 is %s \n", c1)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("err=", err)
		return
	}

}
