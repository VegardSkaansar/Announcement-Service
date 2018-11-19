package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func discord() {
	content := make(map[string]string) // Content for discord

	discordWebhook := "https://discordapp.com/api/webhooks/509465968297115649/BSiX9dBPcDeqegQZ8Q7rRkm8V7Ajgh5RkWkwYFU2uwcXDJxQOkjxgBaHnVCArA0PI9z6"
	content["content"] = "Holla"
	jsonResp, err := json.Marshal(content)
	if err != nil {
		fmt.Println("Error marshaling JSON:")
	}

	_, err = http.Post(discordWebhook, "application/json", bytes.NewBuffer(jsonResp))
	if err != nil {
		fmt.Println("Error making POST request to discord:", err.Error())
	}
}

func clockTrigger() {
	delay := time.Second * 10

	for {

		discord()

		time.Sleep(delay)
	}
}

func main() {
	clockTrigger()
}
