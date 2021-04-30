package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	config "github.com/kamenevn/goapi/config"
	"github.com/kamenevn/goapi/helpers"
	"github.com/kamenevn/goapi/helpers/map"
	"github.com/kamenevn/goapi/models"
	"github.com/kamenevn/goapi/models/handlers"
	"github.com/kamenevn/goapi/modules/request"
	"github.com/kamenevn/goapi/modules/security"
	"net/http"
	"time"
)

/*
Вместо gorilla/mux можно использовать github.com/julienschmidt/httprouter - он быстрее

import (
	//"fmt"
	"encoding/json"
	"net/http"
	"log"
	//"os"
	//"github.com/hhkbp2/go-logging"
	//"github.com/dghubble/sling"
	//"github.com/jinzhu/gorm"

	security "github.com/kamenevn/goapi/modules/security"
	"github.com/kamenevn/goapi/models"
	"github.com/gorilla/mux"
	"time"
	"fmt"
	//"strings"
	"strings"
)
*/

func init() {
	config.AppConfig = config.Get()
	config.DB = config.InitDb("postgres", config.AppConfig)
}

func main() {
	r := mux.NewRouter()

	//r.Schemes("https")

	r.HandleFunc("/", handlers.EmptyHandler)
	r.HandleFunc("/healthcheck", handlers.HealthCheckHandler)

	var bodydata string

	allRoutes, err := models.GetRoutes(config.DB)
	if err == nil {
		//fmt.Print(allRoutes)
		for _, oneRoute := range allRoutes {
			handlerMap := map[string]string {"handler": "DefaultHandler"}
			if len(oneRoute.CustomHandler) > 0 {
				handlerMap["handler"] = oneRoute.CustomHandler
			}

			r.HandleFunc(oneRoute.Input, func(w http.ResponseWriter, r *http.Request) {
				if oneRoute.CheckAccess == true {
					access := security.CheckAccess(oneRoute, w, r)
					if access == false {
						request.ResponseError(w, http.StatusForbidden, "Access denied", nil)
						return
					}
				}

				urlVars := mux.Vars(r)

				switch oneRoute.HttpType {
					case "GET", "DELETE":
						bodydata = r.URL.Query().Encode()
					case "POST", "PUT":
						r.ParseForm()
						bodydata = r.Form.Encode()
					//case "PUT":
					//	_, data := models.Put(values)
					//case "DELETE":
					//	_, data := models.Delete(values)
					default:
						http.Error(w, "Forbidden", 403)
						return
				}

				_map.Call(handlerMap, "handler", w, r, oneRoute, urlVars, bodydata)
			}).Methods(oneRoute.HttpType)
		}
	}

	r.NotFoundHandler = http.HandlerFunc(handlers.EmptyHandler)

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8081",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		helpers.CheckErr(err, "ListenAndServe: ")
	}
}