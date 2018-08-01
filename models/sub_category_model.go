package models

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var subCategoryDB *sql.DB

func init(){
	var err error
	subCategoryDB, err = sql.Open("mysql", "blog:Wzj769397@tcp(123.56.156.172:3306)/blog?parseTime=true")
	if err != nil{
		log.Fatalln(err)
	}
	subCategoryDB.SetMaxIdleConns(20)
	subCategoryDB.SetMaxOpenConns(20)
	if err := subCategoryDB.Ping(); err != nil{
		log.Fatalln(err)
	}
}

func InsertOneSubCategoryRecord(uuid string, title string, main_uuid string,createTime string) (int64, error) {
	rs, err := subCategoryDB.Exec("INSERT INTO sub_category(uuid, title, main_uuid, create_time) VALUES (?, ?, ?, ?)", uuid, title, main_uuid, createTime)
	if err != nil {
		return -1, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}