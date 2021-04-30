package handlers

import (
	"github.com/kamenevn/goapi/helpers/request"
	"github.com/kamenevn/goapi/models"
	requestModel "github.com/kamenevn/goapi/modules/request"
	"net/http"
)

type MainHandler struct{}

func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	requestModel.ResponseError(w, http.StatusNotFound, "Route not found", nil)
	return
}

func (h MainHandler) DefaultHandler(w http.ResponseWriter, r *http.Request, route models.Router, urlvars map[string]string, bodydata string)  {
	outputURL := _request.GetOutputUrl(route.Output, urlvars)
	outputURL = route.Scheme+"://"+route.Domain+outputURL

	_request.DoWithoutAnswer(w, outputURL, route.HttpType, bodydata)
	return
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	requestModel.ResponseSuccess(w, "Health status: ok", nil)
	return
}