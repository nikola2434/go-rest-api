package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func SavePEMKey(filename string, key *rsa.PrivateKey) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file for %s: %v", filename, err)
	}
	defer file.Close()

	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}
	if err := pem.Encode(file, privateKeyPEM); err != nil {
		return fmt.Errorf("failed to write %s to file: %v", filename, err)
	}

	return nil
}

func SavePEMPublicKey(fileName string, key *rsa.PublicKey) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file for %s: %v", fileName, err)
	}
	defer file.Close()

	publicKeyPEM, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return fmt.Errorf("failed to marshal public key: %v", err)
	}

	pemBlock := &pem.Block{Type: "PUBLIC KEY", Bytes: publicKeyPEM}
	if err := pem.Encode(file, pemBlock); err != nil {
		return fmt.Errorf("failed to write %s to file: %v", fileName, err)
	}

	return nil
}

func GenerateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

func main() {
	refreshPrivateKey, refreshPublicKey, err := GenerateRSAKeyPair()
	if err != nil {
		fmt.Printf("Error generating refresh key pair: %v\n", err)
		return
	}

	if err := SavePEMKey("refresh-private-key.pem", refreshPrivateKey); err != nil {
		fmt.Printf("Error saving refresh private key: %v\n", err)
		return
	}

	if err := SavePEMPublicKey("refresh-public-key.pem", refreshPublicKey); err != nil {
		fmt.Printf("Error saving refresh public key: %v\n", err)
		return
	}

	// Генерация и сохранение ключей для access токена
	accessPrivateKey, accessPublicKey, err := GenerateRSAKeyPair()
	if err != nil {
		fmt.Printf("Error generating access key pair: %v\n", err)
		return
	}

	if err := SavePEMKey("access-private-key.pem", accessPrivateKey); err != nil {
		fmt.Printf("Error saving access private key: %v\n", err)
		return
	}

	if err := SavePEMPublicKey("access-public-key.pem", accessPublicKey); err != nil {
		fmt.Printf("Error saving access public key: %v\n", err)
		return
	}

	fmt.Println("Keys generated and saved successfully.")
}
