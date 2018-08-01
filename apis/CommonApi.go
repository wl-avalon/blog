package apis

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"blog/config"
	"log"
	"errors"
)

func Post(businessID string, interfaceID string, params url.Values) (string, error) {

	requestConf, ok := config.RequestConfMap[businessID]
	if !ok{
		log.Fatal("没有找到businessID对应的配置,businessID为" + businessID)

		return "", errors.New("网络请求失败")
	}

	interfaceName, ok := requestConf.InterfaceMap[interfaceID]
	if !ok{
		log.Fatal("没有找到interfaceID对应的配置,interfaceID为" + interfaceID)
		return "", errors.New("网络请求失败")
	}

	requestUrl := "http://" + requestConf.Domain + interfaceName

	resp, err := http.PostForm(requestUrl, params)

	if err != nil {
		log.Fatal("请求失败,error为" + err.Error())
		return "", errors.New("网络请求失败")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("网络请求读取Body时失败,error为" + err.Error())
		return "", errors.New("网络请求失败")
	}
	return string(body), nil
}