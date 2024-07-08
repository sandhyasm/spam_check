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

	userIP, errip := getIP()
	if errip != nil {
		log.Fatal(errip)
		log.Println(errip)
		fmt.Println(errip)
	}

	request := &spamRequest{
		MethodName:     "check_message",
		AuthKey:        "hdfhdhgdhvgdhgb",
		SenderEmail:    "abc@test.com",
		SenderNickname: "Abc Test",
		SenderIp:       userIP,
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

	var response spamResponse
	resErr := json.NewDecoder(res.Body).Decode(&response)
	if resErr != nil {
		log.Fatal("Error while decoding the response")
		log.Fatal(resErr)
		fmt.Println(resErr)
	}

	fmt.Println(resErr)
}

func getIP() (string, error) {
	res, err := http.Get("https://ipify.org/?format=text")
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
