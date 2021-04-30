package models

type Response struct {
	Message  string
	Status   string
	Data   map[string]interface{}
}