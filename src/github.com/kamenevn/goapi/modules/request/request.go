package request

import (
	"encoding/json"
	"github.com/kamenevn/goapi/models"
	"log"
	"net"
	"net/http"
	"strconv"
)

func ResponseError(w http.ResponseWriter, code int, message string, data map[string]interface{}) {
	response := &models.Response{
		Message: message,
		Status: "error",
		Data: data,
	}

	ResponseJSON(w, code, response)
}

func ResponseSuccess(w http.ResponseWriter, message string, data map[string]interface{}) {
	response := &models.Response{
		Message: message,
		Status: "success",
		Data: data,
	}

	ResponseJSON(w, http.StatusOK, response)
}

func ResponseJSON(w http.ResponseWriter, code int, responseModel *models.Response) {
	response, _ := json.Marshal(responseModel)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetClientIdFromHeader(r *http.Request) string {
	return r.Header.Get("CLIENT-ID")
}

func GetClientTimeFromHeader(r *http.Request) int32 {
	clientTime := r.Header.Get("CLIENT-TIME")
	clientTimeInt, err := strconv.Atoi(clientTime)
	if err != nil {
		return 0
	}

	return int32(clientTimeInt)
}

func GetClientHashFromHeader(r *http.Request) string {
	return r.Header.Get("CLIENT-HASH")
}

func GetIPAdress(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Print("userIP: [", r.RemoteAddr, "] is not IP:port")
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		log.Print("userIP: [", r.RemoteAddr, "] is not IP:port")
		return ""
	}

	return userIP.String()
}