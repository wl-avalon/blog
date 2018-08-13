package models

import (
	"database/sql"
	"log"
)

const articleDetailTableName="article_detail"
var articleDetailDB *sql.DB
type ArticleSummary struct {
	Uuid string `json:"uuid" form:"uuid"`
	Title string `json:"title" form:"title"`
	CreateTime string `json:"createTime" form:"create_time"`
	BrowserCount string `json:"browserCount" form:"browser_count"`
}
type ArticleContent struct {
	Uuid string `json:"uuid" form:"uuid"`
	Title string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}
type ArticleCount struct {
	ArticleCount int `json:"listCount" form:"list_count"`
}

func init(){
	var err error
	articleDetailDB, err = sql.Open("mysql", "blog:Wzj769397@tcp(123.56.156.172:3306)/blog?parseTime=true")
	if err != nil{
		log.Fatalln(err)
	}
	articleDetailDB.SetMaxIdleConns(20)
	articleDetailDB.SetMaxOpenConns(20)
	articleDetailDB.SetConnMaxLifetime(0)
	if err := articleDetailDB.Ping(); err != nil{
		log.Fatalln(err)
	}
}

func InsertOneArticleDetailRecord(uuid string, mainCategoryUuid string, subCategoryUuid string, title string, content string, createTime string) (int64, error) {
	rs, err := mainCategoryDB.Exec("INSERT INTO " + articleDetailTableName + "(uuid, main_category_uuid, sub_category_uuid, title, content, create_time) VALUES (?, ?, ?, ?, ?, ?)", uuid, mainCategoryUuid, subCategoryUuid, title, content, createTime)
	if err != nil {
		return -1, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func QueryArticleSummaryListByCategoryUuid(mainUuid string, subUuid string, pageNo int, pageSize int) ([]ArticleSummary, error) {
	offset := (pageNo - 1) * pageSize
	limit  := pageSize

	var rows *sql.Rows
	var err error
	if subUuid == "" && mainUuid == ""{
		rows, err = mainCategoryDB.Query("SELECT uuid, title,browser_count,create_time FROM " + articleDetailTableName + " ORDER BY create_time DESC limit ?,?", offset, limit)
	}else if subUuid == "" {
		rows, err = mainCategoryDB.Query("SELECT uuid, title,browser_count,create_time FROM " + articleDetailTableName + " where main_category_uuid=? ORDER BY create_time DESC limit ?,?", mainUuid, offset, limit)
	}else if mainUuid == "" {
		rows, err = mainCategoryDB.Query("SELECT uuid, title,browser_count,create_time FROM " + articleDetailTableName + " where sub_category_uuid=? ORDER BY create_time DESC limit ?,?", subUuid, offset, limit)
	}else{
		rows, err = mainCategoryDB.Query("SELECT uuid, title,browser_count,create_time FROM " + articleDetailTableName + " where main_category_uuid=? and sub_category_uuid=? ORDER BY create_time DESC limit ?,?", mainUuid, subUuid, offset, limit)
	}
	if err != nil {
		return nil, err
	}

	recordList := make([]ArticleSummary, 0)
	for rows.Next() {
		var articleSummary ArticleSummary
		rows.Scan(&articleSummary.Uuid, &articleSummary.Title, &articleSummary.BrowserCount, &articleSummary.CreateTime)
		recordList = append(recordList, articleSummary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return recordList, nil
}

func QueryArticleDetail(articleUuid string) (ArticleContent, error) {
	var articleDetail ArticleContent
	err := mainCategoryDB.QueryRow("SELECT uuid, title,content FROM " + articleDetailTableName + " where uuid = ?", articleUuid).Scan(&articleDetail.Uuid, &articleDetail.Title, &articleDetail.Content)
	if err != nil {
		return articleDetail, err
	}
	return articleDetail, nil
}

func ModifyArticle(articleUuid string, title string, content string) error {
	_, err := mainCategoryDB.Exec("update " + articleDetailTableName + " set title=?,content=? where uuid=?", title, content, articleUuid)
	if err != nil {
		return err
	}
	return nil
}

func QueryArticleSummaryListCountByCategoryUuid(mainUuid string, subUuid string) (ArticleCount, error) {
	var articleCount ArticleCount
	var err error
	if subUuid == "" && mainUuid == ""{
		err = mainCategoryDB.QueryRow("SELECT count(1) as list_count FROM " + articleDetailTableName + " ORDER BY create_time DESC").Scan(&articleCount.ArticleCount)
	}else if subUuid == "" {
		err = mainCategoryDB.QueryRow("SELECT count(1) as list_count FROM " + articleDetailTableName + " where main_category_uuid=? ORDER BY create_time DESC", mainUuid).Scan(&articleCount.ArticleCount)
	}else if mainUuid == "" {
		err = mainCategoryDB.QueryRow("SELECT count(1) as list_count FROM " + articleDetailTableName + " where sub_category_uuid=? ORDER BY create_time DESC", subUuid).Scan(&articleCount.ArticleCount)
	}else{
		err = mainCategoryDB.QueryRow("SELECT count(1) as list_count FROM " + articleDetailTableName + " where main_category_uuid=? and sub_category_uuid=? ORDER BY create_time DESC", mainUuid, subUuid).Scan(&articleCount.ArticleCount)
	}
	if err != nil {
		return articleCount, err
	}
	return articleCount, nil
}