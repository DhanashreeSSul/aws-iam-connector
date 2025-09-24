package main

import "fmt"

func main() {
	// 1. IAM Connector
	fmt.Println("Fetching IAM Users and Permissions...")
	report, err := listAllUsers()
	if err != nil {
		fmt.Println("IAM Error:", err)
	} else {
		fmt.Println(report)
	}

	// 2. SSH Key Management
	fmt.Println("Generating SSH Keys...")
	demoSSHKeyGeneration()

	// 3. Hooks System
	fmt.Println("Sending Webhook...")
	sendWebhook("https://outlook.office.com/webhook/XXX/YYY/ZZZ", "IAM report + SSH keys generated successfully")
}