package query

import (
	"github.com/gin-gonic/gin"
	"blog/models"
	"net/http"
	"strconv"
)

func GetArticleSummaryList(c *gin.Context){
	mainUuid 		:= c.DefaultQuery("mainUuid", "")
	subUuid			:= c.DefaultQuery("subUuid", "")
	pageNoStr		:= c.DefaultQuery("pageNo", "1")
	pageSizeStr		:= c.DefaultQuery("pageSize", "20")
	pageNo,err		:= strconv.Atoi(pageNoStr)
	if err != nil {
		//fmt.Println(err.Error())
		pageNo = 1
	}
	pageSize,err	:= strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 20
	}

	articleSummaryList, err := models.QueryArticleSummaryListByCategoryUuid(mainUuid, subUuid, pageNo, pageSize)
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

	articleListCount, err := models.QueryArticleSummaryListCountByCategoryUuid(mainUuid, subUuid)
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
		Data: map[string]interface{}{
			"articleList": articleSummaryList,
			"listCount": articleListCount.ArticleCount,
		},
	}
	c.JSON(http.StatusOK, response)
}
