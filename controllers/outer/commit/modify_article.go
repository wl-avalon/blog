package commit

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"blog/models"
)

func ModifyArticle(c *gin.Context){
	uuid := c.DefaultPostForm("uuid", "")
	title := c.DefaultPostForm("title", "")
	content := c.DefaultPostForm("content", "")
	if uuid == "" || content == "" || title == "" {
		response := httpResponse{
			Error: errorStructure{
				ReturnCode: 1,
				ReturnMessage: "uuid,title,content不能为空",
				ReturnUserMessage: "参数错误",
			},
			Data: map[string]interface{}{},
		}
		c.JSON(http.StatusOK, response)
		return
	}

	err := models.ModifyArticle(uuid, title, content)
	if err != nil {
		response := httpResponse{
			Error: errorStructure{
				ReturnCode: 1,
				ReturnMessage: "更新数据库失败," + err.Error(),
				ReturnUserMessage: "更新数据库失败",
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
		Data: map[string]string{},
	}
	c.JSON(http.StatusOK, response)
	return
}
