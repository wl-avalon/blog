package query

import (
	"github.com/gin-gonic/gin"
	"blog/models"
	"net/http"
)

func GetSubCategoryList(c *gin.Context){
	mainUuid := c.DefaultQuery("mainUuid", "0")
	if mainUuid == "0"{
		response := httpResponse{
			Error: errorStructure{
				ReturnCode: 0,
				ReturnMessage: "成功",
				ReturnUserMessage: "成功",
			},
			Data: nil,
		}
		c.JSON(http.StatusOK, response)
		return
	}

	subCategoryRecordList, err := models.QueryAllSubCategoryRecordList(mainUuid)
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
		Data: subCategoryRecordList,
	}
	c.JSON(http.StatusOK, response)
	return
}