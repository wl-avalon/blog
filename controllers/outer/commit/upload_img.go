package commit

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"os"
	"io"
	"blog/apis"
	"math/rand"
	"time"
)

func UploadImg(c *gin.Context){
	file, _ , err := c.Request.FormFile("image")
	var imgPath string
	var uuid string
	for i := 0; i < 10; i++ {
		uuid, err = apis.GetNextUuid()
		if err != nil {
			uuid = strconv.Itoa(rand.Int())
		}
		//imgPath = "/home/saber/webroot/image/blog/" + uuid + ".png"
		imgPath = "/Users/avalonspace/Downloads/" + uuid + ".png"
		_, err = os.Stat(imgPath)
		if err != nil && os.IsNotExist(err) {
			break
		}
		time.Sleep(500 * time.Millisecond)
		continue
	}

	if imgPath == "" || uuid == "" {
		response := httpResponse{
			Error: errorStructure{
				ReturnCode: 1,
				ReturnMessage: "生成图片名失败",
				ReturnUserMessage: "生成图片名失败",
			},
			Data: map[string]string{},
		}
		c.JSON(http.StatusOK, response)
		return
	}
	out, err := os.Create(imgPath)
	if err != nil {
		response := httpResponse{
			Error: errorStructure{
				ReturnCode: 1,
				ReturnMessage: "生成图片失败,error:" + err.Error(),
				ReturnUserMessage: "生成图片失败",
			},
			Data: map[string]string{},
		}
		c.JSON(http.StatusOK, response)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		response := httpResponse{
			Error: errorStructure{
				ReturnCode: 1,
				ReturnMessage: "生成图片失败,error:" + err.Error(),
				ReturnUserMessage: "生成图片失败",
			},
			Data: map[string]string{},
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
			"imgUrl": "https://wl-avalon.com/image/blog/" + uuid + ".png",
		},
	}
	c.JSON(http.StatusOK, response)
	return
}