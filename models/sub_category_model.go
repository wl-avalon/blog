package models

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

const subCategoryTableName="sub_category"
var subCategoryDB *sql.DB

type SubCategoryRecord struct {
	Uuid string `json:"uuid" form:"uuid"`
	Title string `json:"title" form:"title"`
}

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

func InsertOneSubCategoryRecord(uuid string, title string, mainUuid string,createTime string) (int64, error) {
	rs, err := subCategoryDB.Exec("INSERT INTO " + subCategoryTableName + " (uuid, title, main_uuid, create_time) VALUES (?, ?, ?, ?)", uuid, title, mainUuid, createTime)
	if err != nil {
		return -1, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func QueryAllSubCategoryRecordList(mainUuid string) ([]SubCategoryRecord, error) {
	rows, err := mainCategoryDB.Query("SELECT uuid, title FROM " + subCategoryTableName + " WHERE main_uuid=? ORDER BY id", mainUuid)
	if err != nil {
		return nil, err
	}

	recordList := make([]SubCategoryRecord, 0)
	for rows.Next() {
		var subCategoryRecord SubCategoryRecord
		rows.Scan(&subCategoryRecord.Uuid, &subCategoryRecord.Title)
		recordList = append(recordList, subCategoryRecord)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return recordList, nil
}