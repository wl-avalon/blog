package commit

import (
	"github.com/gin-gonic/gin"
	"blog/apis"
	"time"
	"math/rand"
	"strconv"
	"blog/models"
	"net/http"
)

type errorStructure struct{
	ReturnCode int `json:"returnCode"`
	ReturnMessage string `json:"returnMessage"`
	ReturnUserMessage string `json:"returnUserMessage"`
}

type httpResponse struct {
	Error errorStructure `json:"error"`
	Data interface{} `json:"data"`
}

func AddMainCategory(c *gin.Context){
	title := c.DefaultPostForm("title", "")
	uuid, err := apis.GetNextUuid()
	if err != nil {
		uuid = strconv.Itoa(rand.Int())
	}
	createTime := time.Now().Format("2006-01-02 15:04:05")

	lastID, err := models.InsertOneMainCategoryRecord(uuid, title, createTime)
	if err != nil {
		response := httpResponse{
			Error: errorStructure{
				ReturnCode: 1,
				ReturnMessage: "插入数据库失败",
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