# AWS IAM Connector  

This project is a **Go-based connector** for managing AWS IAM users, automating **SSH key management**, and sending **webhook notifications** to external systems.  

It is designed as a learning project for:  
- 🔹 **Integration & Connectors** (AWS IAM, external APIs)  
- 🔹 **SSH Key Management Automation**  
- 🔹 **Hooks/Webhook System** for event notifications  

---

## Project Structure  

```
aws-iam-connector/
│── main.go          # Entry point
│── iam.go           # AWS IAM logic (list users, fetch permissions)
│── ssh_manager.go   # SSH key generation & management
│── hooks.go         # Webhook sender
│── README.md        # Documentation
│── go.mod           # Go module definition
│── go.sum           # Go dependency checksums
│── .gitignore       # Ignore secrets/keys
```

---

## Prerequisites  

1. **Install Go** → [Download Go](https://go.dev/dl/)  

   ```bash
   go version
   ```

2. **AWS Account Setup**  
   - Create an IAM user with programmatic access  
   - Download the `.csv` containing **Access Key** and **Secret Key**  

3. **Configure AWS CLI** (saves credentials locally)  

   ```bash
   aws configure
   ```

   Example inputs:  
   ```
   AWS Access Key ID: <your-access-key>
   AWS Secret Access Key: <your-secret-key>
   Default region: ap-south-1
   ```

4. **Install Dependencies**  

   ```bash
   go mod init github.com/yourusername/aws-iam-connector
   go get github.com/aws/aws-sdk-go-v2/config
   go get github.com/aws/aws-sdk-go-v2/service/iam
   ```

---

## ▶️ Running the Project  

Run all Go files together:  

```bash
go run .
```

---

## SSH Key Management  

- 🔑 Generates a new **RSA SSH key pair (2048-bit)**  
- 💾 Saves the private key locally as `id_rsa`  
- 🔗 Can attach the public key to IAM users (extendable)  

⚠️ **Note:** Never commit `id_rsa` to GitHub. It is ignored in `.gitignore`.  

---

## Webhook System  

The connector can notify external systems (**Slack, Teams, Webhook.site**).  

### Example: Test with Webhook.site  
1. Go to [Webhook.site](https://webhook.site)  
2. Copy your unique webhook URL  
3. Update in `main.go`:  

   ```go
   sendWebhook("https://webhook.site/xxxx-xxxx", "IAM + SSH process complete ✅")
   ```

4. Run the program → Check the Webhook.site page  

---

### Slack Example  

```go
sendWebhook("https://hooks.slack.com/services/XXX/YYY/ZZZ", "IAM update: ✅")
```

---

### Microsoft Teams Example  

```go
sendWebhook("https://outlook.office.com/webhook/XXX/YYY/ZZZ", "IAM update: ✅")
```

---
