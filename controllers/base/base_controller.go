package base

import (
	"github.com/gin-gonic/gin"
	"blog/controllers/outer/commit"
	"net/http"
)

var router *gin.Engine

type RequestParams struct {
	c *gin.Context
}

var outerInterfaceMapFunc  = map[string]func(RequestParams)interface{}{
	"commit/addMainCategory": commit.AddMainCategory,
}

var innerInterfaceMapFunc  = map[string]func(RequestParams)interface{}{
	"commit/addMainCategory": commit.AddMainCategory,
}

func Init(){
	router = gin.Default()
	router.POST("/blog/inner/:actionType/:interfaceName", doInnerHandler)
	router.POST("/blog/outer/:actionType/:interfaceName", doOuterHandler)
}

func Run(){
	router.Run(":8000")
}

func doOuterHandler(c *gin.Context) {
	actionType := c.Param("actionType")
	interfaceName := c.Param("interfaceName")
	url := actionType + "/" + interfaceName

	handleFunc,ok := outerInterfaceMapFunc[url]
	if !ok {
		c.String(http.StatusNotFound, "")
		return
	}

	params := RequestParams{
		c: c,
	}
	str := handleFunc(params)
	c.JSON(http.StatusOK, str)
}

func doInnerHandler(c *gin.Context) {
	actionType := c.Param("actionType")
	interfaceName := c.Param("interfaceName")
	url := actionType + "/" + interfaceName

	handleFunc,ok := innerInterfaceMapFunc[url]
	if !ok {
		c.String(http.StatusNotFound, "")
		return
	}
	params := RequestParams{
		c: c,
	}
	str := handleFunc(params)
	c.JSON(http.StatusOK, str)
}