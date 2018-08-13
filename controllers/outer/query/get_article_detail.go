package query

import (
	"github.com/gin-gonic/gin"
	"blog/models"
	"net/http"
)

func GetArticleDetail(c *gin.Context){
	articleUuid := c.DefaultQuery("articleUuid", "")

	articleDetail, err := models.QueryArticleDetail(articleUuid)
	if err != nil {
		response := httpResponse{
			Error: errorStructure{
				ReturnCode: 1,
				ReturnMessage: "查询数据库失败",
				ReturnUserMessage: "网络繁忙，请稍后再试",
			},
			Data: map[string]interface{}{},
		}
		c.JSON(http.StatusOK, response)
		return
	}

	response := httpResponse{
		Error: errorStructure{
			ReturnCode: 0,
			ReturnMessage: "成功",
			ReturnUserMessage: "成功",
		},
		Data: articleDetail,
	}
	c.JSON(http.StatusOK, response)
}
