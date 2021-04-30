package security

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"github.com/kamenevn/goapi/helpers"
	"github.com/kamenevn/goapi/models"
	"github.com/kamenevn/goapi/modules/request"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func generateHash(clientTime string, clientId string, privateKey string, r *http.Request, w http.ResponseWriter) string {
	r.ParseForm()

	var data string

	switch r.Method {
		case "GET", "DELETE":
			data = r.URL.Query().Encode()
		case "POST", "PUT":
			r.ParseForm()
			data = r.Form.Encode()
		default:
			request.ResponseError(w, http.StatusBadRequest, "Method not allowed", nil)
	}

	notHashedString := clientTime+clientId+data
	key := []byte(privateKey)

	sig := hmac.New(sha512.New, key)
	sig.Write([]byte(notHashedString))

	return hex.EncodeToString(sig.Sum(nil))
}

func checkHash(generatedHash string, clientHash string) bool {
	clientHashDecode, _ := helpers.Base64Decode(clientHash)
	return strings.EqualFold(generatedHash, clientHashDecode)
}

func CheckAccess(router models.Router, w http.ResponseWriter, r *http.Request) bool {
	clientIp := request.GetIPAdress(r)
	clientId := request.GetClientIdFromHeader(r)
	clientTime := request.GetClientTimeFromHeader(r)
	clientHash := request.GetClientHashFromHeader(r)

	if clientIp == "" || clientId == "" || clientTime == 0 || clientHash == "" {
		return false
	}

	nowTimestamp := int32(time.Now().Unix())
	if ((nowTimestamp - clientTime) >= 300) {
		return false
	}

	merchant, err := models.CheckAccess(clientId, clientIp, router.Id)
	if (err != nil) {
		return false
	}

	generatedHash := generateHash(strconv.FormatInt(int64(clientTime), 10), clientId, merchant.PrivateKey, r, w)
	result := checkHash(generatedHash, clientHash)

	return result
}