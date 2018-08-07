package models

import (
	"database/sql"
	"log"
)

const articleDetailTableName="article_detail"
var articleDetailDB *sql.DB
type ArticleDetail struct {
	Uuid string `json:"uuid" form:"uuid"`
	Title string `json:"title" form:"title"`
}

func init(){
	var err error
	articleDetailDB, err = sql.Open("mysql", "blog:Wzj769397@tcp(123.56.156.172:3306)/blog?parseTime=true")
	if err != nil{
		log.Fatalln(err)
	}
	articleDetailDB.SetMaxIdleConns(20)
	articleDetailDB.SetMaxOpenConns(20)
	if err := articleDetailDB.Ping(); err != nil{
		log.Fatalln(err)
	}
}

func InsertOneArticleDetailRecord(uuid string, mainCategoryUuid string, subCategoryUuid string, content string, createTime string) (int64, error) {
	rs, err := mainCategoryDB.Exec("INSERT INTO " + articleDetailTableName + "(uuid, main_category_uuid, sub_category_uuid, content, create_time) VALUES (?, ?, ?, ?, ?)", uuid, mainCategoryUuid, subCategoryUuid, content, createTime)
	if err != nil {
		return -1, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

//func QueryAllMainCategoryRecordList() ([]MainCategoryRecord, error) {
//	rows, err := mainCategoryDB.Query("SELECT uuid, title FROM " + mainCategoryTableName + " ORDER BY id")
//	if err != nil {
//		return nil, err
//	}
//
//	recordList := make([]MainCategoryRecord, 0)
//	for rows.Next() {
//		var mainCategoryRecord MainCategoryRecord
//		rows.Scan(&mainCategoryRecord.Uuid, &mainCategoryRecord.Title)
//		recordList = append(recordList, mainCategoryRecord)
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//	return recordList, nil
//}