package _request

import (
	"encoding/json"
	"github.com/kamenevn/goapi/models"
	requestModel "github.com/kamenevn/goapi/modules/request"
	"github.com/valyala/fasthttp"
	"net/http"
	"regexp"
	"strings"
)

func GetOutputUrl(outputURL string, urlvars map[string]string) string {
	var resultUrl []string
	urlPaths := strings.Split(outputURL, "/")

	// Заменить на http://elliot.land/post/go-replace-string-with-regular-expression-callback
	for _, onePath := range urlPaths {
		if strings.Contains(onePath, "{") {
			onePath = strings.Replace(onePath, "{", "", -1)
			onePath = strings.Replace(onePath, "}", "", -1)
			splitPath := strings.Split(onePath, ":")
			varName := splitPath[0]
			i, ok := urlvars[varName]
			if splitPath[1] != "" && ok {
				delete(urlvars, varName)
				matched, _ := regexp.MatchString(splitPath[1], i)
				if matched {
					onePath = i
				}
			}
		}
		resultUrl = append(resultUrl, onePath)
	}

	return strings.Join(resultUrl, "/")
}

func DoWithoutAnswer(w http.ResponseWriter, url string, httpMethod string, bodydata string)  {
	response, err := DoWithAnswer(url, httpMethod, bodydata)

	if err != nil || string(response) == "" {
		message := "Ошибка сервера"
		if err != nil {
			message = err.Error()
		}
		requestModel.ResponseError(w, http.StatusBadRequest, message, nil)
		return
	}

	responseModel := models.Response{}
	json.Unmarshal([]byte(response), &responseModel)

	requestModel.ResponseJSON(w, http.StatusBadRequest, &responseModel)
	return
}

func DoWithAnswer(url string, httpMethod string, bodydata string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(httpMethod)

	switch httpMethod {
		case "GET", "DELETE":
			url = url+"?"+bodydata
		case "POST", "PUT":
			req.SetBodyString(bodydata)
	}

	var bodyBytes []byte

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err == nil {
		bodyBytes = resp.Body()
	}

	return bodyBytes, err
}