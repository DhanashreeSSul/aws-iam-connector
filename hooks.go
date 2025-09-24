package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func sendWebhook(url string, message string) {
	payload := []byte(fmt.Sprintf(`{"event": "%s"}`, message))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending webhook:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Webhook sent, status:", resp.Status)
}
