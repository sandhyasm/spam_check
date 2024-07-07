package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type spamRequest struct {
	MethodName     string `json:"method_name"`
	AuthKey        string `json:"auth_key"`
	SenderEmail    string `json:"sender_email"`
	SenderNickname string `json:"sender_nickname"`
	SenderIp       string `json:"sender_ip"`
	JsOn           int    `json:"js_on"`
	SubmitTime     int    `json:"submit_time"`
}

type spamResponse struct {
	Data map[string]interface{} `json:"data"`
}

func main() {

	/*
		data='{"method_name":"check_message",
		"auth_key":"your_acccess_key",
		"sender_email":"stop_email@example.com",
		"sender_nickname":"John Doe",
		"sender_ip":"127.0.0.1",
		"js_on":1,"submit_time":15}' https://moderate.cleantalk.org/api2.0
	*/

	request := &spamRequest{
		MethodName:     "check_message",
		AuthKey:        "nysumygepuvetud",
		SenderEmail:    "abc@test.com",
		SenderNickname: "Abc Test",
		SenderIp:       "120.18.17.10",
		JsOn:           1,
		SubmitTime:     15,
	}

	url := "https://moderate.cleantalk.org/api2.0"
	jsonReq, err := json.Marshal(request)
	if err != nil {
		log.Fatal("Issue while marshalling the json call")
		log.Fatal(err)
		fmt.Println(err)
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatal("Issue while calling the api")
		log.Fatal(err)
		fmt.Println(err)
	}

	fmt.Println(res)

	fmt.Println("Init checks")
}
