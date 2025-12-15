package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	yandexdisk "github.com/tigusigalpa/yandex-disk-go"
)

func main() {
	accessToken := os.Getenv("YANDEX_DISK_TOKEN")
	if accessToken == "" {
		log.Fatal("YANDEX_DISK_TOKEN environment variable is not set")
	}

	client := yandexdisk.NewClient(accessToken)

	tempFile, err := os.CreateTemp("", "yandex-disk-test-*.txt")
	if err != nil {
		log.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	testContent := "Hello from Yandex Disk Go SDK!\nThis is a test file."
	if _, err := tempFile.WriteString(testContent); err != nil {
		log.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	fmt.Printf("Created temporary file: %s\n", tempFile.Name())

	remotePath := "/disk/go-sdk-test/" + filepath.Base(tempFile.Name())
	fmt.Printf("Uploading to: %s\n", remotePath)

	result, err := client.UploadFile(tempFile.Name(), remotePath, true)
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	fmt.Printf("Upload Status: %d\n", result.Status)
	fmt.Printf("Upload Success: %v\n", result.Success)

	if result.Success {
		fmt.Println("\n=== Getting uploaded file metadata ===")
		meta, err := client.GetMeta(remotePath, nil)
		if err != nil {
			log.Printf("Failed to get metadata: %v", err)
		} else {
			fmt.Printf("Name: %s\n", meta.Name)
			fmt.Printf("Size: %d bytes\n", meta.Size)
			fmt.Printf("MD5: %s\n", meta.MD5)
			fmt.Printf("Created: %s\n", meta.Created)
		}

		fmt.Println("\n=== Downloading file back ===")
		downloadPath := filepath.Join(os.TempDir(), "downloaded-"+filepath.Base(tempFile.Name()))
		err = client.DownloadFile(remotePath, downloadPath)
		if err != nil {
			log.Printf("Failed to download file: %v", err)
		} else {
			fmt.Printf("Downloaded to: %s\n", downloadPath)
			
			content, err := os.ReadFile(downloadPath)
			if err != nil {
				log.Printf("Failed to read downloaded file: %v", err)
			} else {
				fmt.Printf("Content: %s\n", string(content))
			}
			
			os.Remove(downloadPath)
		}

		fmt.Println("\n=== Cleaning up: deleting remote file ===")
		err = client.Delete(remotePath, true)
		if err != nil {
			log.Printf("Failed to delete file: %v", err)
		} else {
			fmt.Println("File deleted successfully")
		}
	}

	fmt.Println("\n=== Upload example completed ===")
}
