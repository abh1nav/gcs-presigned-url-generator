package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

func main() {
	projectID := flag.String("project", "", "Google Cloud Project ID")
	bucketName := flag.String("bucket", "", "GCS Bucket Name")
	objectKey := flag.String("key", "", "Object Key")
	credentialsFile := flag.String("credentials", "", "Path to service account key JSON file")

	flag.Parse()

	if *projectID == "" || *bucketName == "" || *objectKey == "" || *credentialsFile == "" {
		log.Fatal("All flags --project, --bucket, --key and --credentialsFile must be provided.")
	}

	url, err := createURL(*bucketName, *objectKey, *credentialsFile)
	if err != nil {
		log.Fatalf("Failed to generate signed URL: %v", err)
	}

	fmt.Println("Signed URL:")
	fmt.Println(url)
}

func createURL(bucketName, objectKey, credentialsFile string) (string, error) {
	credentialsJSON, err := os.ReadFile(credentialsFile)
	if err != nil {
		return "", fmt.Errorf("Unable to read credentials file: %v", err)
	}

	var credentials struct {
		ClientEmail string `json:"client_email"`
		PrivateKey  string `json:"private_key"`
	}
	if err := json.Unmarshal(credentialsJSON, &credentials); err != nil {
		return "", fmt.Errorf("Unable to parse credentials JSON: %v", err)
	}

	expires := time.Now().Add(12 * time.Hour)
	opts := &storage.SignedURLOptions{
		GoogleAccessID: credentials.ClientEmail,
		PrivateKey:     []byte(credentials.PrivateKey),
		Method:         "POST",
		Expires:        expires,
		Headers:        []string{"x-goog-resumable:start"},
	}

	url, err := storage.SignedURL(bucketName, objectKey, opts)
	if err != nil {
		return "", err
	}

	return url, nil
}
