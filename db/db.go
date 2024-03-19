package db

// _ "github.com/mattn/go-sqlite3"
import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB
func InitDB() {
	var err error
	// 打开（或创建）一个名为 "api.db" 的 SQLite 数据库文件
	DB, err = sql.Open("sqlite", "api.db")
 
	if err != nil {
		panic(fmt.Errorf("打开或创建数据库失败:%w", err))
	}

	// 最多可以同时打开10个连接
	DB.SetMaxOpenConns(10)
	// 最大空闲连接数（如果没有人用保持5个连接）
	DB.SetMaxIdleConns(5)
 
	createTables()
}


func createTables() {
	/**	
		db.Exec 用于执行不返回结果集的 SQL 命令（如 INSERT, UPDATE, DELETE, CREATE TABLE 等）
		而 db.Query 用于执行返回结果集的 SQL 命令（如 SELECT）
		- CREATE TABLE IF NOT EXISTS events 如果events表不存在，则创建一个
		- id INTEGER PRIMARY KEY AUTOINCREMENT 定义一个名为id的列，其数据类型为整数（INTEGER）
			该列被指定为主键（PRIMARY KEY）这意味着它的值必须是唯一的，并且不能为null
			AUTOINCREMENT关键字表示每当向表中插入新行时，id列的值都会自动增加。
			这通常用于确保每个行都一个唯一的标识符
		- name TEXT NOT NULL定义了一个名为name的列，其数据类型为文本（TEXT）。NOT NULL约束表示
			该列不能包含NULL值，即每个users表中的行都必须有一个name
		- dateTime DATETIME NOT NULL 定义一个名为dateTime的列，其数据类型为日期（DATETIME）	
		- user_id INTEGER 定义一个名为userID的列，其数据类型为整数（INTEGER）

	*/
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			datetime DATETIME NOT NULL,
			user_id INTEGER
		)
	`

	_,err := DB.Exec(createEventsTable)


	if err != nil {
		panic(fmt.Errorf("创建表events失败:%w", err))
	}

}