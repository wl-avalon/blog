package models

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var mainCategoryDB *sql.DB

func init(){
	var err error
	mainCategoryDB, err = sql.Open("mysql", "blog:Wzj769397@tcp(123.56.156.172:3306)/blog?parseTime=true")
	if err != nil{
		log.Fatalln(err)
	}
	mainCategoryDB.SetMaxIdleConns(20)
	mainCategoryDB.SetMaxOpenConns(20)
	if err := mainCategoryDB.Ping(); err != nil{
		log.Fatalln(err)
	}
}

func InsertOneMainCategoryRecord(uuid string, title string, createTime string) (int64, error) {
	rs, err := mainCategoryDB.Exec("INSERT INTO main_category(uuid, title, create_time) VALUES (?, ?, ?)", uuid, title, createTime)
	if err != nil {
		return -1, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}