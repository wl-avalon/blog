package commit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddArticle(c *gin.Context){
	c.String(http.StatusOK, "asdasdasdasd")
}
