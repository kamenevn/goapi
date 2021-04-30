package handlers

import (
	"net/http"
)

func (h MainHandler) AllHandler(w http.ResponseWriter, r *http.Request, params ... interface{}) {
	/*
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response := &models.Response{
		MESSAGE: "",
		STATUS: "test",
	}

	fmt.Println(params)

	responseJson, _ := json.Marshal(response)
	fmt.Println(string(responseJson))
	*/
	return
}
