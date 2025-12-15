package main

import (
	"fmt"
	"log"
	"os"

	yandexdisk "github.com/tigusigalpa/yandex-disk-go"
)

func main() {
	accessToken := os.Getenv("YANDEX_DISK_TOKEN")
	if accessToken == "" {
		log.Fatal("YANDEX_DISK_TOKEN environment variable is not set")
	}

	client := yandexdisk.NewClient(accessToken)

	diskInfo, err := client.GetCapacity()
	if err != nil {
		log.Fatalf("Failed to get disk info: %v", err)
	}

	fmt.Println("=== Disk Information ===")
	fmt.Printf("Total Space: %d bytes (%.2f GB)\n", diskInfo.TotalSpace, float64(diskInfo.TotalSpace)/1024/1024/1024)
	fmt.Printf("Used Space: %d bytes (%.2f GB)\n", diskInfo.UsedSpace, float64(diskInfo.UsedSpace)/1024/1024/1024)
	fmt.Printf("Free Space: %d bytes (%.2f GB)\n", diskInfo.GetFreeSpace(), float64(diskInfo.GetFreeSpace())/1024/1024/1024)
	fmt.Printf("Usage: %.2f%%\n", diskInfo.GetUsagePercentage())
	fmt.Printf("Trash Size: %d bytes (%.2f MB)\n", diskInfo.TrashSize, float64(diskInfo.TrashSize)/1024/1024)
	fmt.Printf("Is Paid: %v\n", diskInfo.IsPaid)
	fmt.Printf("User: %s (%s)\n", diskInfo.User.DisplayName, diskInfo.User.Login)

	fmt.Println("\n=== Recent Uploads ===")
	recentUploads, err := client.GetRecentUploads(5, 0)
	if err != nil {
		log.Printf("Failed to get recent uploads: %v", err)
	} else {
		for i, file := range recentUploads.Items {
			fmt.Printf("%d. %s (%d bytes) - %s\n", i+1, file.Name, file.Size, file.Modified)
		}
	}

	fmt.Println("\n=== Creating Test Folder ===")
	testFolder := "/disk/go-sdk-test"
	folder, err := client.CreateFolder(testFolder)
	if err != nil {
		log.Printf("Failed to create folder (may already exist): %v", err)
	} else {
		fmt.Printf("Created folder: %s\n", folder.Path)
	}

	fmt.Println("\n=== Getting Folder Metadata ===")
	meta, err := client.GetMeta(testFolder, nil)
	if err != nil {
		log.Printf("Failed to get metadata: %v", err)
	} else {
		fmt.Printf("Name: %s\n", meta.Name)
		fmt.Printf("Type: %s\n", meta.Type)
		fmt.Printf("Created: %s\n", meta.Created)
		fmt.Printf("Modified: %s\n", meta.Modified)
	}

	fmt.Println("\n=== Example completed successfully ===")
}
