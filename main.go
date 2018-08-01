package main

import (
	"github.com/gin-gonic/gin"
	"blog/controllers/outer/commit"
)

func main(){
	router := gin.Default()
	router.POST("/blog/outer/commit/addArticle", commit.AddArticle)
	router.POST("/blog/outer/commit/addMainCategory", commit.AddMainCategory)
	router.POST("/blog/outer/commit/addSubCategory", commit.AddSubCategory)
	router.Run(":8000")
}