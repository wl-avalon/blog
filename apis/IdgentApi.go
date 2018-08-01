package apis

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type getNextUuidResponse struct {
	RequestType string `json:"type"`
	NextId int64
}

func GetNextUuid() (string,error) {
	response, err := Post("idgent", "nextID", nil)
	if err != nil {
		fmt.Println("请求失败：16 error：" + err.Error())
		return "", err
	}

	var formatResponse getNextUuidResponse
	err = json.Unmarshal([]byte(response), &formatResponse)
	if err != nil {
		fmt.Println("请求失败：23 error：" + err.Error())
		return "", err
	}

	nextID :=  strconv.FormatInt(formatResponse.NextId,10)
	return nextID, nil
}