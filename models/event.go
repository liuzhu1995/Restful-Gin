package models

import (
	"fmt"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string`binding:"required"`
	Description string`binding:"required"`
	Location    string`binding:"required"`
	DateTime    time.Time`binding:"required"`
	UserID int64
}


func (e *Event) Save() error {
	// 添加到数据库

	//  准备插入语句  VALUES (?, ?, ?, ?, ?) 参数占位符提高性能并减少SQL注入的风险
	query := `
		INSERT INTO events (name, description, location, dateTime, user_id) 
		VALUES (?, ?, ?, ?, ?)`
	// DB.Prepare 方法用于准备SQL语句，并返回一个 *sql.Stmt 对象，该对象代表了一个预编译的SQL语句	
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("预编译SQL语句失败:%w",err)
	}
	defer stmt.Close()
 
	// 插入数据
	result,err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return fmt.Errorf("插入数据失败:%w", err)
	}
	// 获取最后一个插入操作生成的自增 ID
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("获取自增id失败:%w", err)
	}
	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	// for rows.Next() 循环用于迭代查询结果集中的每一行
	for rows.Next() {
		// 对于每一行创建一个新的结构体实例event
		var event Event
		// 使用 rows.Scan() 方法将该行的数据填充到结构体的字段中
		 err := rows.Scan(
			&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime,&event.UserID,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}	 
	return events, nil
}

func GetEventByID(id int64)(*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("预编译SQL语句失败:%v", err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	var event Event
	// 使用 Scan 方法，将每行数据填充到 event 变量中
	err = row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &event, nil
}

func (e Event)UpdateEvent() error {
	fmt.Printf("%+v\n", e)
	query := `UPDATE events 
		SET name = ?,description = ?,location = ?,dateTime = ?
		WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	if err != nil {
		return fmt.Errorf("更新数据失败:%w", err)
	}
	
	return nil
}


func (e Event)DeleteEvent() error {
	query := "DELETE FROM events where id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("SQL语句预编译失败:%v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}
