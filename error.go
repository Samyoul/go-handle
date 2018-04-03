package handlers

import (
	"log"
	"fmt"
)

type ErrorHandler struct {}

func (e ErrorHandler) Handle(err error, message string, v ...interface{}) {
	formattedMessage := fmt.Sprintf(message, v...) + fmt.Sprint(" - Error : ", err)
	e.fatalf(err, formattedMessage)
}

func (e ErrorHandler) fatalf(err error, message string) {
	if err != nil {
		log.Fatalf(message)
	}
}
