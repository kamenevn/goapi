package helpers

import (
	"encoding/base64"
	"fmt"
	"log"
)

var WritoToConsole = true

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
		panic(err.Error())
	}
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) (string, bool) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", true
	}
	return string(data), false
}

func PrintToConsole(msg string) {
	if WritoToConsole == true {
		fmt.Println(msg)
	}
}