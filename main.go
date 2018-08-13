package main

import (
	"github.com/gin-gonic/gin"
	"blog/controllers/outer/commit"
	"blog/controllers/outer/query"
	"strings"
	"fmt"
)

func main(){
	router := gin.Default()
	router.Use(Cors())

	router.POST("/blog/outer/commit/addArticle", commit.AddArticle)
	router.POST("/blog/outer/commit/addMainCategory", commit.AddMainCategory)
	router.POST("/blog/outer/commit/addSubCategory", commit.AddSubCategory)
	router.POST("/blog/outer/commit/modifyArticle", commit.ModifyArticle)
	router.POST("/blog/outer/commit/uploadImg", commit.UploadImg)

	router.GET("/blog/outer/query/getMainCategoryList", query.GetMainCategoryList)
	router.GET("/blog/outer/query/getSubCategoryList", query.GetSubCategoryList)
	router.GET("/blog/outer/query/getArticleSummaryList", query.GetArticleSummaryList)
	router.GET("/blog/outer/query/getArticleDetail", query.GetArticleDetail)
	router.Run(":8000")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			//下面的都是乱添加的-_-~
			// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET")
			// c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			// c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}

		c.Next()
	}
}