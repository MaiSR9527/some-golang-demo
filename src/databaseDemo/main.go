package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/16 13:44
 *@version v1.0
 */
func main() {
	db, err := sql.Open("mysql", "root:mai1208142397@tcp(localhost:3306)/test")
	db.Ping()
	defer func() {
		if db != nil {

			db.Close()
		}
	}()
	if err != nil {
		fmt.Println("数据库连接异常", err.Error())
		return
	}
	stmt, preErr := db.Prepare("insert into payment values (default ,?,?,?,?)")
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if preErr != nil {
		fmt.Println("预处理失败", preErr.Error())
		return
	}
	result, stmtErr := stmt.Exec("红米Note", 22000.0, time.Now().Format("2006-01-02 15:04:05"), "0pPKHjWprnVxGH7dEsAoXX2YQvU")
	if stmtErr != nil {
		fmt.Println("sql执行失败", stmtErr.Error())
		return
	}
	count, rowErr := result.RowsAffected()
	if rowErr != nil {
		fmt.Println("获取结果失败")
		return
	}
	if count > 0 {
		fmt.Println("新增成功")
	} else {
		fmt.Println("新增失败")
	}

}
