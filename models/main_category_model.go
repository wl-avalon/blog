package models

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

const mainCategoryTableName="main_category"
var mainCategoryDB *sql.DB
type MainCategoryRecord struct {
	Uuid string `json:"uuid" form:"uuid"`
	Title string `json:"title" form:"title"`
}

func init(){
	var err error
	mainCategoryDB, err = sql.Open("mysql", "blog:Wzj769397@tcp(123.56.156.172:3306)/blog?parseTime=true")
	if err != nil{
		log.Fatalln(err)
	}
	mainCategoryDB.SetMaxIdleConns(20)
	mainCategoryDB.SetMaxOpenConns(20)
	mainCategoryDB.SetConnMaxLifetime(0)
	if err := mainCategoryDB.Ping(); err != nil{
		log.Fatalln(err)
	}
}

func InsertOneMainCategoryRecord(uuid string, title string, createTime string) (int64, error) {
	rs, err := mainCategoryDB.Exec("INSERT INTO " + mainCategoryTableName + "(uuid, title, create_time) VALUES (?, ?, ?)", uuid, title, createTime)
	if err != nil {
		return -1, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func QueryAllMainCategoryRecordList() ([]MainCategoryRecord, error) {
	rows, err := mainCategoryDB.Query("SELECT uuid, title FROM " + mainCategoryTableName + " ORDER BY id")
	if err != nil {
		return nil, err
	}

	recordList := make([]MainCategoryRecord, 0)
	for rows.Next() {
		var mainCategoryRecord MainCategoryRecord
		rows.Scan(&mainCategoryRecord.Uuid, &mainCategoryRecord.Title)
		recordList = append(recordList, mainCategoryRecord)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return recordList, nil
}