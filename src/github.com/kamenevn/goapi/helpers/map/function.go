package _map

import (
	"github.com/kamenevn/goapi/models"
	"github.com/kamenevn/goapi/models/handlers"
	"net/http"
	"reflect"
)

func Call(m map[string] string, name string, w http.ResponseWriter, r *http.Request, route models.Router, urlvars map[string]string, bodydata string) (result []reflect.Value, err error) {
	mHandler := handlers.MainHandler{}
	functionByName := reflect.ValueOf(&mHandler).MethodByName(m[name])

	result = functionByName.Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r), reflect.ValueOf(route), reflect.ValueOf(urlvars), reflect.ValueOf(bodydata)})
	return
}