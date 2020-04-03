/**
 * @Author: dragon
 * @Description:
 * @File:  uesr
 * @Version: 1.0.0
 * @Date: 2020/4/2 下午2:50
 */

package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func PrintUsers() {
	db, err := sql.Open("mysql", "root:yn666666@tcp(localhost:3306)/blogserver?charset=utf8")
	defer db.Close()
	if err != nil {
		fmt.Println("数据库打开失败", err)
		return
	}

	stmt, err := db.Prepare("select * from user")
	defer stmt.Close()
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query()
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		var sex int
		rows.Scan(&id, &name, &sex)
		fmt.Println(id, name, sex)
	}

}
