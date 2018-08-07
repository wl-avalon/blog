package query

import (
	"github.com/gin-gonic/gin"
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

func GetMainCategoryList(c *gin.Context){
	mainCategoryRecordList, err := models.QueryAllMainCategoryRecordList()
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
	}

	response := httpResponse{
		Error: errorStructure{
			ReturnCode: 0,
			ReturnMessage: "成功",
			ReturnUserMessage: "成功",
		},
		Data: mainCategoryRecordList,
	}
	c.JSON(http.StatusOK, response)
}