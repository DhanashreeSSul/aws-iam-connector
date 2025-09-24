package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/ssh"
)

func generateSSHKeyPair() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privatePEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	})

	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	publicKey := ssh.MarshalAuthorizedKey(pub)

	return string(privatePEM), string(publicKey), nil
}

func demoSSHKeyGeneration() {
	priv, pub, err := generateSSHKeyPair()
	if err != nil {
		fmt.Println("Error generating SSH key:", err)
		return
	}
	fmt.Println("SSH Key Pair Generated")
	fmt.Println("Private Key (first 100 chars):", priv[:100], "...")
	fmt.Println("Public Key:", pub)
}
