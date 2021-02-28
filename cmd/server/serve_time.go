package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CurrentTime struct {
	Hours   string `json:"Hours"`
	Minutes string `json:"Minutes"`
	Seconds string `json:"Seconds"`
}

func serveTime(writer http.ResponseWriter, request *http.Request) {
	log.Println("Endpoint Hit: serveTime")

	hours, minutes, seconds := time.Now().Clock()
	currentTime := CurrentTime{
		Hours:   strconv.Itoa(hours),
		Minutes: strconv.Itoa(minutes),
		Seconds: strconv.Itoa(seconds),
	}

	_ = json.NewEncoder(writer).Encode(currentTime)
}
