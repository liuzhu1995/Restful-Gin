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
	UserID int
}


func (e *Event) Save() error {
	// 添加到数据库

	//  准备插入语句  VALUES (?, ?, ?, ?, ?) 参数占位符提高性能并减少SQL注入的风险
	query := `
		INSERT INTO events (name, description, location, dateTime, user_id) 
		VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("准备插入语句失败:%w",err)
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
	for rows.Next() {
		var event Event
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
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return &Event{}, err
	}

	return &event, nil
}