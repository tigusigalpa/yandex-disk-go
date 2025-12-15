# Yandex Disk Go SDK

<div align="center">
  <img src="https://github.com/user-attachments/assets/e68c499e-7bc7-4dd1-bb91-319ffa138ab4" alt="Yandex Disk Go SDK" style="max-width: 100%; height: auto;">
</div>

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Yandex Disk API](https://img.shields.io/badge/API-Yandex%20Disk%20API-orange.svg)](https://yandex.ru/dev/disk-api/doc/ru/)

**🌐 Language:** [Русский](README.md) | English

A comprehensive Go SDK for integration with [Yandex Disk API](https://yandex.ru/dev/disk-api/doc/ru/). This library provides a clean, idiomatic Go interface for managing files and folders on Yandex Disk with full coverage of the official API.

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

## 🚀 Installation

```bash
go get github.com/tigusigalpa/yandex-disk-go
```

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

## 📖 Usage Examples

See [README.md](README.md) for detailed usage examples in Russian, or refer to the test files for code examples.

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

| Category              | Implemented | Total | Percent  |
|-----------------------|-------------|-------|----------|
| Disk Information      | ✅ 1         | 1     | 100%     |
| File Operations       | ✅ 8         | 8     | 100%     |
| Public Resources      | ✅ 8         | 8     | 100%     |
| Trash Management      | ✅ 3         | 3     | 100%     |
| Metadata              | ✅ 2         | 2     | 100%     |
| Admin Methods         | ✅ 3         | 3     | 100%     |
| Operations            | ✅ 1         | 1     | 100%     |
| **Total**             | ✅ 26        | 26    | **100%** |

## 🤝 Contributing

Contributions are welcome! Feel free to submit a Pull Request.

## 📄 License

This package is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 🔗 Links

- [Official Yandex Disk API Documentation](https://yandex.ru/dev/disk-api/doc/ru/)
- [OAuth Documentation](https://yandex.ru/dev/id/doc/ru/)
- [GitHub Repository](https://github.com/tigusigalpa/yandex-disk-go)

## 📞 Support

For issues, questions, or contributions:

- Create an issue on [GitHub](https://github.com/tigusigalpa/yandex-disk-go/issues)
- Check the [official documentation](https://yandex.ru/dev/disk-api/doc/ru/)
- Review the [troubleshooting guide](https://yandex.ru/dev/disk-api/doc/ru/concepts/troubleshooting)

---

**Made with ❤️ for the Go community**
