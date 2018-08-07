package commit

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"blog/apis"
	"strconv"
	"time"
	"blog/models"
	"math/rand"
)

func AddArticle(c *gin.Context){
	mainUuid := c.DefaultPostForm("mainUuid", "-1")
	subUuid := c.DefaultPostForm("subUuid", "-1")
	content := c.DefaultPostForm("content", "")
	if mainUuid == "-1" || subUuid == "-1" || content == "" {
		response := httpResponse{
			Error: errorStructure{
				ReturnCode: 1,
				ReturnMessage: "mainUuid,subUuid,content不能为空",
				ReturnUserMessage: "参数错误",
			},
			Data: map[string]interface{}{},
		}
		c.JSON(http.StatusOK, response)
		return
	}
	uuid, err := apis.GetNextUuid()
	if err != nil {
		uuid = strconv.Itoa(rand.Int())
	}
	createTime := time.Now().Format("2006-01-02 15:04:05")

	lastID, err := models.InsertOneArticleDetailRecord(uuid, mainUuid, subUuid, content,createTime)
	if err != nil {
		response := httpResponse{
			Error: errorStructure{
				ReturnCode: 1,
				ReturnMessage: "插入数据库失败," + err.Error(),
				ReturnUserMessage: "插入数据库失败",
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
		Data: map[string]string{
			"lastID": strconv.FormatInt(lastID, 10),
		},
	}
	c.JSON(http.StatusOK, response)
	return
}