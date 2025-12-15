<div align="center">

# 🚀 Yandex Disk Go SDK

<img src="https://github.com/user-attachments/assets/e68c499e-7bc7-4dd1-bb91-319ffa138ab4" alt="Yandex Disk Go SDK" width="600">

### Powerful, modern, and easy-to-use Go SDK for Yandex.Disk

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/tigusigalpa/yandex-disk-go?style=for-the-badge)](https://goreportcard.com/report/github.com/tigusigalpa/yandex-disk-go)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/tigusigalpa/yandex-disk-go)

[![Tests](https://img.shields.io/github/actions/workflow/status/tigusigalpa/yandex-disk-go/test.yml?branch=main&label=Tests&style=for-the-badge)](https://github.com/tigusigalpa/yandex-disk-go/actions)
[![Coverage](https://img.shields.io/codecov/c/github/tigusigalpa/yandex-disk-go?style=for-the-badge&logo=codecov)](https://codecov.io/gh/tigusigalpa/yandex-disk-go)
[![Yandex Disk API](https://img.shields.io/badge/API-Yandex%20Disk-FFCC00?style=for-the-badge&logo=yandex)](https://yandex.ru/dev/disk-api/doc/ru/)

**🌐 Language:** [Русский](README.md) | [English](#)

---

</div>

## ✨ Features

<table>
<tr>
<td width="50%">

### 🎯 Complete API Coverage
- ✅ **26/26 methods** implemented
- 📁 File and folder management
- 🔗 Public links and sharing
- 🗑️ Trash management
- 👥 Admin functions
- 🔄 Async operations

</td>
<td width="50%">

### 💎 Code Quality
- 🧪 Full test coverage
- 📖 Comprehensive documentation
- 🎨 Idiomatic Go code
- ⚡ High performance
- 🛡️ Type safety
- 🔍 Detailed error handling

</td>
</tr>
</table>

## 🎬 Quick Start

### Installation

```bash
go get github.com/tigusigalpa/yandex-disk-go
```

### Your First Program

```go
package main

import (
    "fmt"
    "log"
    
    yandexdisk "github.com/tigusigalpa/yandex-disk-go"
)

func main() {
    // Create client
    client := yandexdisk.NewClient("your_oauth_token")
    
    // Get disk info
    diskInfo, err := client.GetCapacity()
    if err != nil {
        log.Fatal(err)
    }
    
    // Display stats
    fmt.Printf("💾 Used: %.2f GB of %.2f GB\n", 
        float64(diskInfo.UsedSpace)/1e9,
        float64(diskInfo.TotalSpace)/1e9)
    fmt.Printf("📊 Usage: %.1f%%\n", diskInfo.GetUsagePercentage())
    
    // Upload file
    result, _ := client.UploadFile("local.txt", "/disk/remote.txt", true)
    if result.Success {
        fmt.Println("✅ File uploaded successfully!")
    }
}
```

## 🎨 Capabilities

### 📤 Upload and Download Files
```go
// Upload file
result, err := client.UploadFile("photo.jpg", "/disk/photos/photo.jpg", true)

// Download file
err = client.DownloadFile("/disk/document.pdf", "./local/document.pdf")

// Upload from internet
op, err := client.UploadFromURL("https://example.com/file.zip", "/disk/file.zip", false)
```

### 🔗 Public Links
```go
// Publish file and get link
resource, err := client.Publish("/disk/presentation.pptx")
fmt.Printf("Share: %s\n", resource.PublicURL)

// Download someone's public file
err = client.DownloadPublicResource("https://yadi.sk/d/...", "./downloaded.pdf", nil)

// Save public file to your disk
resource, err = client.SavePublicResource("https://yadi.sk/d/...", nil, stringPtr("/disk/saved/"))
```

### 📊 Space Management
```go
// Get disk statistics
diskInfo, err := client.GetCapacity()
fmt.Printf("📦 Free: %.2f GB\n", float64(diskInfo.GetFreeSpace())/1e9)
fmt.Printf("🗑️ In trash: %.2f MB\n", float64(diskInfo.TrashSize)/1e6)

// Clear trash
err = client.ClearTrash(nil)
```

### 🔍 Search and Filtering
```go
// Get all files
files, err := client.GetAllFiles(100, 0)

// Recently uploaded
recent, err := client.GetRecentUploads(20, 0)

// Published files
published, err := client.GetRecentPublished(10, 0)
```

### 🏢 For Organizations
```go
// Admin functions for managing user resources
resources, err := client.GetPublicResourcesOwnedByUser("user-id", "org-id", 20, 0)
err = client.UnpublishUserResource("public-key", "user-id", "org-id")
```

---

## 📋 API Reference

| Method                               | Endpoint                                           | Documentation                                                                                        | Description                                   |
|--------------------------------------|----------------------------------------------------|-----------------------------------------------------------------------------------------------------|-----------------------------------------------|
| `GetAuthorizationURL()`              | -                                                  | [OAuth Guide](https://yandex.ru/dev/disk-api/doc/ru/concepts/quickstart)                            | Generate OAuth authorization URL              |
| `GetCapacity()`                      | `GET /`                                            | [Disk Info](https://yandex.ru/dev/disk-api/doc/ru/reference/capacity)                               | Get disk information                          |
| `GetMeta()`                          | `GET /resources`                                   | [Metadata](https://yandex.ru/dev/disk-api/doc/ru/reference/meta)                                    | Get resource metadata                         |
| `AddMeta()`                          | `PATCH /resources`                                 | [Add Metadata](https://yandex.ru/dev/disk-api/doc/ru/reference/meta-add)                            | Add custom metadata                           |
| `GetAllFiles()`                      | `GET /resources/files`                             | [All Files](https://yandex.ru/dev/disk-api/doc/ru/reference/all-files)                              | Get flat list of all files                    |
| `GetRecentUploads()`                 | `GET /resources/last-uploaded`                     | [Recent Uploads](https://yandex.ru/dev/disk-api/doc/ru/reference/recent-upload)                     | Get recently uploaded files                   |
| `GetRecentPublished()`               | `GET /resources/public`                            | [Published Files](https://yandex.ru/dev/disk-api/doc/ru/reference/recent-public)                    | Get recently published files                  |
| `CreateFolder()`                     | `PUT /resources`                                   | [Create Folder](https://yandex.ru/dev/disk-api/doc/ru/reference/create-folder)                      | Create folder                                 |
| `UploadFile()`                       | `GET /resources/upload`                            | [Upload File](https://yandex.ru/dev/disk-api/doc/ru/reference/upload)                               | Upload file                                   |
| `UploadFromURL()`                    | `POST /resources/upload`                           | [Upload from URL](https://yandex.ru/dev/disk-api/doc/ru/reference/upload-ext)                       | Upload file from internet                     |
| `DownloadFile()`                     | `GET /resources/download`                          | [Download File](https://yandex.ru/dev/disk-api/doc/ru/reference/content)                            | Download file                                 |
| `Copy()`                             | `POST /resources/copy`                             | [Copy Resource](https://yandex.ru/dev/disk-api/doc/ru/reference/copy)                               | Copy file/folder                              |
| `Move()`                             | `POST /resources/move`                             | [Move Resource](https://yandex.ru/dev/disk-api/doc/ru/reference/move)                               | Move file/folder                              |
| `Delete()`                           | `DELETE /resources`                                | [Delete Resource](https://yandex.ru/dev/disk-api/doc/ru/reference/delete)                           | Delete file/folder                            |
| `Publish()`                          | `PUT /resources/publish`                           | [Publish Resource](https://yandex.ru/dev/disk-api/doc/ru/reference/publish)                         | Publish resource                              |
| `Unpublish()`                        | `PUT /resources/unpublish`                         | [Unpublish Resource](https://yandex.ru/dev/disk-api/doc/ru/reference/unpublish)                     | Unpublish resource                            |
| `GetAvailablePublicSettings()`       | `GET /public/resources/public-settings/available`  | [Available Settings](https://yandex.ru/dev/disk-api/doc/ru/reference/public-settings-get-available) | Get available public settings                 |
| `GetPublicSettings()`                | `GET /public/resources/public-settings`            | [Public Settings](https://yandex.ru/dev/disk-api/doc/ru/reference/public-settings-get)              | Get public settings for resource              |
| `ChangePublicSettings()`             | `PUT /resources/public`                            | [Change Settings](https://yandex.ru/dev/disk-api/doc/ru/reference/public-settings-change)           | Change public settings                        |
| `GetPublicResourceMeta()`            | `GET /public/resources`                            | [Public Metadata](https://yandex.ru/dev/disk-api/doc/ru/reference/public)                           | Get public resource metadata                  |
| `DownloadPublicResource()`           | `GET /public/resources/download`                   | [Download Public](https://yandex.ru/dev/disk-api/doc/ru/reference/public)                           | Download public resource                      |
| `SavePublicResource()`               | `POST /public/resources/save`                      | [Save Public Resource](https://yandex.ru/dev/disk-api/doc/ru/reference/public)                      | Save public resource                          |
| `GetTrash()`                         | `GET /trash/resources`                             | [Trash List](https://yandex.ru/dev/disk-api/doc/ru/reference/trash-delete)                          | Get trash contents                            |
| `RestoreFromTrash()`                 | `PUT /trash/resources/restore`                     | [Restore from Trash](https://yandex.ru/dev/disk-api/doc/ru/reference/trash-restore)                 | Restore from trash                            |
| `ClearTrash()`                       | `DELETE /trash/resources`                          | [Clear Trash](https://yandex.ru/dev/disk-api/doc/ru/reference/trash-delete)                         | Clear trash                                   |
| `GetOperationStatus()`               | `GET /operations/{id}`                             | [Operation Status](https://yandex.ru/dev/disk-api/doc/ru/reference/operations)                      | Get operation status                          |
| `GetPublicResourcesOwnedByUser()`    | `GET /public/resources/admin/public-resources`     | [Owned Resources](https://yandex.ru/dev/disk-api/doc/ru/reference/public-owned-by-user)             | Admin: user's public resources                |
| `GetPublicResourcesAccessedByUser()` | `GET /public/resources/admin/accessible-resources` | [Accessible Resources](https://yandex.ru/dev/disk-api/doc/ru/reference/public-accessed-by-user)     | Admin: resources accessible by user           |
| `UnpublishUserResource()`            | `PUT /public/resources/admin/unpublish`            | [Admin Unpublish](https://yandex.ru/dev/disk-api/doc/ru/reference/unpublish-admin-phash)            | Admin: unpublish resource                     |

## 📦 Installation

```bash
go get github.com/tigusigalpa/yandex-disk-go
```

<details>
<summary>📚 Requirements</summary>

- Go 1.21 or higher
- Active Yandex account
- OAuth token for API access

</details>

## 🔐 Getting OAuth Token

To work with Yandex Disk API, you need to obtain an OAuth token. Follow these steps:

### 1. Create Application on Yandex OAuth

1. Log in to your Yandex OAuth account: https://oauth.yandex.ru/
2. Click the "+ Create" button
3. In the popup "What application do you want to create?", select "For API access or debugging" and click "Go to creation"
4. Fill out the form:
   - **Service name**: Enter your application name
   - **Contact email**: Your contact email
   - **Data access**: Select required permissions:
     - `cloud_api:disk.write` — Write anywhere on Disk
     - `cloud_api:disk.read` — Read entire Disk
     - `cloud_api:disk.app_folder` — Access to application folder on Disk
     - `cloud_api:disk.info` — Access to Disk information

After creating the application, you will be shown:
- **ClientID** — needed to obtain OAuth token
- **Client secret** — not needed for Yandex Disk

### 2. Generate Authorization Link

```go
package main

import (
	"fmt"
	yandexdisk "github.com/tigusigalpa/yandex-disk-go"
)

func main() {
	clientID := "your_client_id_from_app_settings"
	authURL := yandexdisk.GetAuthorizationURL(clientID)
	
	fmt.Println("Go to this URL for authorization:")
	fmt.Println(authURL)
	// Output: https://oauth.yandex.ru/authorize?response_type=token&client_id=<ClientID>
}
```

### 3. Get Token

1. Go to the generated link
2. Log in to your Yandex account (if not already logged in)
3. Allow access for your application
4. After confirmation, you will be redirected to a page with the token in the URL
5. Copy the token from the page, something like `y0__xCD2tUFGKDjOyD2-Myl...`

### 4. Use Token

```go
package main

import (
	"fmt"
	"log"
	yandexdisk "github.com/tigusigalpa/yandex-disk-go"
)

func main() {
	accessToken := "your_oauth_token"
	client := yandexdisk.NewClient(accessToken)
	
	diskInfo, err := client.GetCapacity()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Used: %d bytes\n", diskInfo.UsedSpace)
	fmt.Printf("Total: %d bytes\n", diskInfo.TotalSpace)
	fmt.Printf("Free: %d bytes\n", diskInfo.GetFreeSpace())
}
```

## 💡 Usage Examples

> 💡 **Tip**: All examples can be found in the [`examples/`](examples/) folder

### 📊 Disk Information

```go
diskInfo, err := client.GetCapacity()
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Used: %d bytes\n", diskInfo.UsedSpace)
fmt.Printf("Total: %d bytes\n", diskInfo.TotalSpace)
fmt.Printf("Free: %d bytes\n", diskInfo.GetFreeSpace())
fmt.Printf("Usage: %.2f%%\n", diskInfo.GetUsagePercentage())
fmt.Printf("Trash: %d bytes\n", diskInfo.TrashSize)
fmt.Printf("Paid: %v\n", diskInfo.IsPaid)
```

### 📁 File Operations

<details>
<summary>📤 Upload File</summary>

```go
result, err := client.UploadFile(
	"/local/path/file.txt",
	"/disk/MyFolder/file.txt",
	true, // overwrite existing file
)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Status: %d\n", result.Status)
fmt.Printf("Success: %v\n", result.Success)
```

</details>

<details>
<summary>📥 Download File</summary>

```go
err := client.DownloadFile(
	"/disk/MyFile.txt",
	"/local/path/downloaded.txt",
)
if err != nil {
	log.Fatal(err)
}

fmt.Println("Download successful")
```

</details>

<details>
<summary>📋 Copy File</summary>

```go
resource, err := client.Copy(
	"/disk/original.txt",
	"/disk/copy.txt",
	true, // overwrite if exists
)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Copied to: %s\n", resource.Path)
```

</details>

<details>
<summary>🚚 Move File</summary>

```go
resource, err := client.Move(
	"/disk/old-location/file.txt",
	"/disk/new-location/file.txt",
	true, // overwrite if exists
)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Moved to: %s\n", resource.Path)
```

</details>

<details>
<summary>🗑️ Delete File</summary>

```go
err := client.Delete(
	"/disk/file.txt",
	true, // permanent deletion (false = to trash)
)
if err != nil {
	log.Fatal(err)
}

fmt.Println("Deleted")
```

</details>

For more examples, see the [Russian README](README.md) or check the [`examples/`](examples/) directory.

## 🔧 Error Handling

The SDK provides comprehensive error handling:

```go
result, err := client.UploadFile("/path/to/file.txt", "/disk/file.txt", true)
if err != nil {
	if apiErr, ok := err.(*yandexdisk.APIError); ok {
		fmt.Printf("API Error: %s\n", apiErr.Error())
		fmt.Printf("Status Code: %d\n", apiErr.StatusCode)
	} else {
		fmt.Printf("Error: %s\n", err.Error())
	}
	return
}
```

## 📊 API Coverage

<div align="center">

### 🎯 100% Official API Coverage

</div>

| Category              | Implemented | Total | Percent  | Status |
|-----------------------|-------------|-------|----------|--------|
| 💾 Disk Information      | 1         | 1     | 100%     | ✅ |
| 📁 File Operations       | 8         | 8     | 100%     | ✅ |
| 🔗 Public Resources      | 8         | 8     | 100%     | ✅ |
| 🗑️ Trash Management      | 3         | 3     | 100%     | ✅ |
| 🏷️ Metadata              | 2         | 2     | 100%     | ✅ |
| 👥 Admin Methods         | 3         | 3     | 100%     | ✅ |
| ⚙️ Operations            | 1         | 1     | 100%     | ✅ |
| **🎉 Total**             | **26**        | **26**    | **100%** | **✅** |

<div align="center">

**All Yandex Disk API methods are implemented and tested!**

</div>

## 🤝 Contributing

<div align="center">

### We welcome your contributions! 🎉

</div>

Contributions are welcome! Here's how you can help:

- 🐛 **Found a bug?** [Create an issue](https://github.com/tigusigalpa/yandex-disk-go/issues/new)
- 💡 **Have an idea?** [Suggest an improvement](https://github.com/tigusigalpa/yandex-disk-go/issues/new)
- 🔧 **Want to help?** [Submit a Pull Request](https://github.com/tigusigalpa/yandex-disk-go/pulls)
- 📖 **Improve documentation** - any help is appreciated!

Read [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## 📄 License

This package is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 🔗 Useful Links

<table>
<tr>
<td>

### 📚 Documentation
- [Yandex Disk API](https://yandex.ru/dev/disk-api/doc/ru/)
- [OAuth Guide](https://yandex.ru/dev/id/doc/ru/)
- [GoDoc](https://pkg.go.dev/github.com/tigusigalpa/yandex-disk-go)

</td>
<td>

### 🛠️ Development
- [GitHub Repository](https://github.com/tigusigalpa/yandex-disk-go)
- [Issue Tracker](https://github.com/tigusigalpa/yandex-disk-go/issues)
- [Changelog](https://github.com/tigusigalpa/yandex-disk-go/releases)

</td>
</tr>
</table>

## 📞 Support

For issues, questions, or contributions:

- Create an issue on [GitHub](https://github.com/tigusigalpa/yandex-disk-go/issues)
- Check the [official documentation](https://yandex.ru/dev/disk-api/doc/ru/)
- Review the [troubleshooting guide](https://yandex.ru/dev/disk-api/doc/ru/concepts/troubleshooting)

---

<div align="center">

### ⭐ Like the project?

**Give it a star on GitHub!**

[![GitHub stars](https://img.shields.io/github/stars/tigusigalpa/yandex-disk-go?style=social)](https://github.com/tigusigalpa/yandex-disk-go/stargazers)

---

**Made with ❤️ for the Go community**

*Author: [Igor Sazonov](https://github.com/tigusigalpa) | License: MIT*

</div>
