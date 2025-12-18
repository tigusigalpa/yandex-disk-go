<div align="center">

# 🚀 Yandex Disk Go SDK

![Yandex Disk Go SDK](https://github.com/user-attachments/assets/192f1442-d23f-4edb-b389-8a8dbc1fa4fd)

### Мощный, современный и простой в использовании Go SDK для Яндекс.Диска

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/tigusigalpa/yandex-disk-go?style=for-the-badge)](https://goreportcard.com/report/github.com/tigusigalpa/yandex-disk-go)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/tigusigalpa/yandex-disk-go)

[![Tests](https://img.shields.io/github/actions/workflow/status/tigusigalpa/yandex-disk-go/test.yml?branch=main&label=Tests&style=for-the-badge)](https://github.com/tigusigalpa/yandex-disk-go/actions)
[![Coverage](https://img.shields.io/codecov/c/github/tigusigalpa/yandex-disk-go?style=for-the-badge&logo=codecov)](https://codecov.io/gh/tigusigalpa/yandex-disk-go)
[![Yandex Disk API](https://img.shields.io/badge/API-Yandex%20Disk-FFCC00?style=for-the-badge&logo=yandex)](https://yandex.ru/dev/disk-api/doc/ru/)

**🌐 Language:** Русский | [English](README-en.md)

---

</div>

## ✨ Особенности

<table>
<tr>
<td width="50%">

### 🎯 Полное покрытие API

- ✅ **26/26 методов** реализовано
- 📁 Управление файлами и папками
- 🔗 Публичные ссылки и шаринг
- 🗑️ Работа с корзиной
- 👥 Административные функции
- 🔄 Асинхронные операции

</td>
<td width="50%">

### 💎 Качество кода

- 🧪 Полное покрытие тестами
- 📖 Подробная документация
- 🎨 Идиоматичный Go код
- ⚡ Высокая производительность
- 🛡️ Типобезопасность
- 🔍 Детальная обработка ошибок

</td>
</tr>
</table>

## 🎬 Быстрый старт

### Установка

```bash
go get github.com/tigusigalpa/yandex-disk-go
```

### Первая программа

```go
package main

import (
    "fmt"
    "log"
    
    yandexdisk "github.com/tigusigalpa/yandex-disk-go"
)

func main() {
    // Создаём клиент
    client := yandexdisk.NewClient("ваш_oauth_токен")
    
    // Получаем информацию о диске
    diskInfo, err := client.GetCapacity()
    if err != nil {
        log.Fatal(err)
    }
    
    // Выводим статистику
    fmt.Printf("💾 Использовано: %.2f GB из %.2f GB\n", 
        float64(diskInfo.UsedSpace)/1e9,
        float64(diskInfo.TotalSpace)/1e9)
    fmt.Printf("📊 Заполнено: %.1f%%\n", diskInfo.GetUsagePercentage())
    
    // Загружаем файл
    result, _ := client.UploadFile("local.txt", "/disk/remote.txt", true)
    if result.Success {
        fmt.Println("✅ Файл успешно загружен!")
    }
}
```

## 🎨 Возможности

### 📤 Загрузка и скачивание файлов

```go
// Загрузка файла с прогресс-баром (концептуально)
result, err := client.UploadFile("photo.jpg", "/disk/photos/photo.jpg", true)

// Скачивание файла
err = client.DownloadFile("/disk/document.pdf", "./local/document.pdf")

// Загрузка из интернета
op, err := client.UploadFromURL("https://example.com/file.zip", "/disk/file.zip", false)
```

### � Публичные ссылки

```go
// Опубликовать файл и получить ссылку
resource, err := client.Publish("/disk/presentation.pptx")
fmt.Printf("Поделитесь: %s\n", resource.PublicURL)

// Скачать чужой публичный файл
err = client.DownloadPublicResource("https://yadi.sk/d/...", "./downloaded.pdf", nil)

// Сохранить публичный файл на свой диск
resource, err = client.SavePublicResource("https://yadi.sk/d/...", nil, stringPtr("/disk/saved/"))
```

### 📊 Управление пространством

```go
// Получить статистику диска
diskInfo, err := client.GetCapacity()
fmt.Printf("📦 Свободно: %.2f GB\n", float64(diskInfo.GetFreeSpace())/1e9)
fmt.Printf("🗑️ В корзине: %.2f MB\n", float64(diskInfo.TrashSize)/1e6)

// Очистить корзину
err = client.ClearTrash(nil)
```

### 🔍 Поиск и фильтрация

```go
// Получить все файлы
files, err := client.GetAllFiles(100, 0)

// Недавно загруженные
recent, err := client.GetRecentUploads(20, 0)

// Опубликованные файлы
published, err := client.GetRecentPublished(10, 0)
```

### 🏢 Для организаций

```go
// Административные функции для управления ресурсами пользователей
resources, err := client.GetPublicResourcesOwnedByUser("user-id", "org-id", 20, 0)
err = client.UnpublishUserResource("public-key", "user-id", "org-id")
```

---

## �📋 Справочник API

| Метод                                | Эндпоинт                                           | Документация                                                                                        | Описание                                      |
|--------------------------------------|----------------------------------------------------|-----------------------------------------------------------------------------------------------------|-----------------------------------------------|
| `GetAuthorizationURL()`              | -                                                  | [OAuth Guide](https://yandex.ru/dev/disk-api/doc/ru/concepts/quickstart)                            | Генерация URL авторизации OAuth               |
| `GetCapacity()`                      | `GET /`                                            | [Disk Info](https://yandex.ru/dev/disk-api/doc/ru/reference/capacity)                               | Получение информации о диске                  |
| `GetMeta()`                          | `GET /resources`                                   | [Metadata](https://yandex.ru/dev/disk-api/doc/ru/reference/meta)                                    | Получение метаданных ресурса                  |
| `AddMeta()`                          | `PATCH /resources`                                 | [Add Metadata](https://yandex.ru/dev/disk-api/doc/ru/reference/meta-add)                            | Добавление пользовательских метаданных        |
| `GetAllFiles()`                      | `GET /resources/files`                             | [All Files](https://yandex.ru/dev/disk-api/doc/ru/reference/all-files)                              | Получение плоского списка всех файлов         |
| `GetRecentUploads()`                 | `GET /resources/last-uploaded`                     | [Recent Uploads](https://yandex.ru/dev/disk-api/doc/ru/reference/recent-upload)                     | Получение недавно загруженных файлов          |
| `GetRecentPublished()`               | `GET /resources/public`                            | [Published Files](https://yandex.ru/dev/disk-api/doc/ru/reference/recent-public)                    | Получение недавно опубликованных файлов       |
| `CreateFolder()`                     | `PUT /resources`                                   | [Create Folder](https://yandex.ru/dev/disk-api/doc/ru/reference/create-folder)                      | Создание папки                                |
| `UploadFile()`                       | `GET /resources/upload`                            | [Upload File](https://yandex.ru/dev/disk-api/doc/ru/reference/upload)                               | Загрузка файла                                |
| `UploadFromURL()`                    | `POST /resources/upload`                           | [Upload from URL](https://yandex.ru/dev/disk-api/doc/ru/reference/upload-ext)                       | Загрузка файла из интернета                   |
| `DownloadFile()`                     | `GET /resources/download`                          | [Download File](https://yandex.ru/dev/disk-api/doc/ru/reference/content)                            | Скачивание файла                              |
| `Copy()`                             | `POST /resources/copy`                             | [Copy Resource](https://yandex.ru/dev/disk-api/doc/ru/reference/copy)                               | Копирование файла/папки                       |
| `Move()`                             | `POST /resources/move`                             | [Move Resource](https://yandex.ru/dev/disk-api/doc/ru/reference/move)                               | Перемещение файла/папки                       |
| `Delete()`                           | `DELETE /resources`                                | [Delete Resource](https://yandex.ru/dev/disk-api/doc/ru/reference/delete)                           | Удаление файла/папки                          |
| `Publish()`                          | `PUT /resources/publish`                           | [Publish Resource](https://yandex.ru/dev/disk-api/doc/ru/reference/publish)                         | Публикация ресурса                            |
| `Unpublish()`                        | `PUT /resources/unpublish`                         | [Unpublish Resource](https://yandex.ru/dev/disk-api/doc/ru/reference/unpublish)                     | Отмена публикации ресурса                     |
| `GetAvailablePublicSettings()`       | `GET /public/resources/public-settings/available`  | [Available Settings](https://yandex.ru/dev/disk-api/doc/ru/reference/public-settings-get-available) | Получение доступных публичных настроек        |
| `GetPublicSettings()`                | `GET /public/resources/public-settings`            | [Public Settings](https://yandex.ru/dev/disk-api/doc/ru/reference/public-settings-get)              | Получение публичных настроек ресурса          |
| `ChangePublicSettings()`             | `PUT /resources/public`                            | [Change Settings](https://yandex.ru/dev/disk-api/doc/ru/reference/public-settings-change)           | Изменение публичных настроек                  |
| `GetPublicResourceMeta()`            | `GET /public/resources`                            | [Public Metadata](https://yandex.ru/dev/disk-api/doc/ru/reference/public)                           | Получение метаданных публичного ресурса       |
| `DownloadPublicResource()`           | `GET /public/resources/download`                   | [Download Public](https://yandex.ru/dev/disk-api/doc/ru/reference/public)                           | Скачивание публичного ресурса                 |
| `SavePublicResource()`               | `POST /public/resources/save`                      | [Save Public Resource](https://yandex.ru/dev/disk-api/doc/ru/reference/public)                      | Сохранение публичного ресурса                 |
| `GetTrash()`                         | `GET /trash/resources`                             | [Trash List](https://yandex.ru/dev/disk-api/doc/ru/reference/trash-delete)                          | Получение содержимого корзины                 |
| `RestoreFromTrash()`                 | `PUT /trash/resources/restore`                     | [Restore from Trash](https://yandex.ru/dev/disk-api/doc/ru/reference/trash-restore)                 | Восстановление из корзины                     |
| `ClearTrash()`                       | `DELETE /trash/resources`                          | [Clear Trash](https://yandex.ru/dev/disk-api/doc/ru/reference/trash-delete)                         | Очистка корзины                               |
| `GetOperationStatus()`               | `GET /operations/{id}`                             | [Operation Status](https://yandex.ru/dev/disk-api/doc/ru/reference/operations)                      | Получение статуса операции                    |
| `GetPublicResourcesOwnedByUser()`    | `GET /public/resources/admin/public-resources`     | [Owned Resources](https://yandex.ru/dev/disk-api/doc/ru/reference/public-owned-by-user)             | Администратор: публичные ресурсы пользователя |
| `GetPublicResourcesAccessedByUser()` | `GET /public/resources/admin/accessible-resources` | [Accessible Resources](https://yandex.ru/dev/disk-api/doc/ru/reference/public-accessed-by-user)     | Администратор: доступные пользователю ресурсы |
| `UnpublishUserResource()`            | `PUT /public/resources/admin/unpublish`            | [Admin Unpublish](https://yandex.ru/dev/disk-api/doc/ru/reference/unpublish-admin-phash)            | Администратор: отмена публикации ресурса      |

## � Установка

```bash
go get github.com/tigusigalpa/yandex-disk-go
```

<details>
<summary>📚 Требования</summary>

- Go 1.21 или выше
- Активный аккаунт Яндекс
- OAuth токен для доступа к API

</details>

## 🔐 Получение OAuth-токена

Для работы с Yandex Disk API необходимо получить OAuth-токен. Следуйте этим шагам:

### 1. Создание приложения на Яндекс OAuth

1. Зайдите под своей учётной записью на Яндекс OAuth: https://oauth.yandex.ru/
2. Нажмите на кнопку "+ Создать"
3. Во всплывающем окне "Какое приложение хотите создать?" укажите "Для доступа к API или отладки" и нажмите "Перейти к
   созданию"
4. Заполните форму:
    - **Название сервиса**: Укажите название вашего приложения
    - **Почта для связи**: Ваш контактный email
    - **Доступ к данным**: Выберите необходимые права:
        - `cloud_api:disk.write` — Запись в любом месте на Диске
        - `cloud_api:disk.read` — Чтение всего Диска
        - `cloud_api:disk.app_folder` — Доступ к папке приложения на Диске
        - `cloud_api:disk.info` — Доступ к информации о Диске

После создания приложения вам будут показаны:

- **ClientID** — понадобится для получения OAuth-токена
- **Client secret** — для работы с Яндекс Диском он не понадобится

### 2. Формирование ссылки авторизации

```go
package main

import (
	"fmt"
	yandexdisk "github.com/tigusigalpa/yandex-disk-go"
)

func main() {
	clientID := "ваш_client_id_из_настроек_приложения"
	authURL := yandexdisk.GetAuthorizationURL(clientID)
	
	fmt.Println("Перейдите по ссылке для авторизации:")
	fmt.Println(authURL)
	// Вывод: https://oauth.yandex.ru/authorize?response_type=token&client_id=<ClientID>
}
```

### 3. Получение токена

1. Перейдите по сгенерированной ссылке
2. Авторизуйтесь в своём аккаунте Яндекса (если еще не авторизованы)
3. Разрешите доступ для вашего приложения
4. После подтверждения вы будете перенаправлены на страницу с токеном в URL
5. Скопируйте токен на странице, что-то вроде `y0__xCD2tUFGKDjOyD2-Myl...`

### 4. Использование токена

```go
package main

import (
	"fmt"
	"log"
	yandexdisk "github.com/tigusigalpa/yandex-disk-go"
)

func main() {
	accessToken := "ваш_oauth_токен"
	client := yandexdisk.NewClient(accessToken)
	
	diskInfo, err := client.GetCapacity()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Использовано: %d байт\n", diskInfo.UsedSpace)
	fmt.Printf("Всего: %d байт\n", diskInfo.TotalSpace)
	fmt.Printf("Свободно: %d байт\n", diskInfo.GetFreeSpace())
}
```

## � Примеры использования

> 💡 **Совет**: Все примеры можно найти в папке [`examples/`](examples/)

### 📊 Информация о диске

```go
diskInfo, err := client.GetCapacity()
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Использовано: %d байт\n", diskInfo.UsedSpace)
fmt.Printf("Всего: %d байт\n", diskInfo.TotalSpace)
fmt.Printf("Свободно: %d байт\n", diskInfo.GetFreeSpace())
fmt.Printf("Использование: %.2f%%\n", diskInfo.GetUsagePercentage())
fmt.Printf("Корзина: %d байт\n", diskInfo.TrashSize)
fmt.Printf("Платный: %v\n", diskInfo.IsPaid)
```

### Работа с ресурсами

```go
resource, err := client.GetMeta("/disk/MyFile.txt", nil)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Имя: %s\n", resource.Name)
fmt.Printf("Тип: %s\n", resource.Type)
fmt.Printf("Размер: %d байт\n", resource.Size)
fmt.Printf("Создан: %s\n", resource.Created)
fmt.Printf("Изменён: %s\n", resource.Modified)
fmt.Printf("MIME тип: %s\n", resource.MimeType)
fmt.Printf("MD5: %s\n", resource.MD5)
fmt.Printf("SHA256: %s\n", resource.SHA256)
```

### 📁 Операции с файлами

<details>
<summary>📤 Загрузка файла</summary>

#### Простая загрузка

```go
result, err := client.UploadFile(
	"/local/path/file.txt",
	"/disk/MyFolder/file.txt",
	true, // перезаписать существующий файл
)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Статус: %d\n", result.Status)
fmt.Printf("Успешно: %v\n", result.Success)
```

</details>

<details>
<summary>📥 Скачивание файла</summary>

```go
err := client.DownloadFile(
	"/disk/MyFile.txt",
	"/local/path/downloaded.txt",
)
if err != nil {
	log.Fatal(err)
}

fmt.Println("Скачивание успешно")
```

</details>

<details>
<summary>📋 Копирование файла</summary>

```go
resource, err := client.Copy(
	"/disk/original.txt",
	"/disk/copy.txt",
	true, // перезаписать, если существует
)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Скопировано в: %s\n", resource.Path)
```

</details>

<details>
<summary>🚚 Перемещение файла</summary>

```go
resource, err := client.Move(
	"/disk/old-location/file.txt",
	"/disk/new-location/file.txt",
	true, // перезаписать, если существует
)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Перемещено в: %s\n", resource.Path)
```

</details>

<details>
<summary>🗑️ Удаление файла</summary>

```go
err := client.Delete(
	"/disk/file.txt",
	true, // окончательное удаление (false = в корзину)
)
if err != nil {
	log.Fatal(err)
}

fmt.Println("Удалено")
```

</details>

### 📂 Операции с папками

```go
resource, err := client.CreateFolder("/disk/MyNewFolder")
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Папка создана: %s\n", resource.Path)

client.CreateFolder("/disk/Projects")
client.CreateFolder("/disk/Projects/WebDev")
client.CreateFolder("/disk/Projects/WebDev/Site1")
```

### Список файлов и папок

```go
allFiles, err := client.GetAllFiles(100, 0)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Всего файлов: %d\n", allFiles.Total)

for _, file := range allFiles.Items {
	fmt.Printf("- %s (%s)\n", file.Name, file.Type)
}

recent, err := client.GetRecentUploads(10, 0)
if err != nil {
	log.Fatal(err)
}
for _, file := range recent.Items {
	fmt.Printf("- %s\n", file.Name)
}

dirMeta, err := client.GetMeta("/disk/MyFolder", nil)
if err != nil {
	log.Fatal(err)
}

if dirMeta.IsDir() {
	for _, item := range dirMeta.GetItems() {
		fmt.Printf("- %s (%d байт)\n", item.Name, item.Size)
	}
	fmt.Printf("Всего элементов: %d\n", dirMeta.GetTotalItems())
}
```

### Управление метаданными

```go
resource, err := client.AddMeta("/disk/file.txt", map[string]interface{}{
	"description": "Моё пользовательское описание",
	"author":      "John Doe",
	"version":     "1.0.0",
})
if err != nil {
	log.Fatal(err)
}

meta, err := client.GetMeta("/disk/file.txt", nil)
if err != nil {
	log.Fatal(err)
}

fmt.Println(meta.CustomProperties)
```

### Управление публичным доступом

#### Публикация ресурса

```go
resource, err := client.Publish("/disk/document.pdf")
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Публичный URL: %s\n", resource.PublicURL)
```

#### Получение опубликованных файлов

```go
published, err := client.GetRecentPublished(10, 0)
if err != nil {
	log.Fatal(err)
}

for _, file := range published.Items {
	if file.IsPublished() {
		fmt.Printf("- %s\n", file.Name)
		fmt.Printf("  URL: %s\n", file.PublicURL)
		fmt.Printf("  Ключ: %s\n", file.PublicKey)
	}
}
```

#### Отмена публикации ресурса

```go
_, err := client.Unpublish("/disk/document.pdf")
if err != nil {
	log.Fatal(err)
}
```

#### Работа с публичными ресурсами

```go
publicMeta, err := client.GetPublicResourceMeta("https://yadi.sk/d/abc123...", nil)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Публичный файл: %s\n", publicMeta.Name)

err = client.DownloadPublicResource("https://yadi.sk/d/abc123...", "/local/downloaded.pdf", nil)
if err != nil {
	log.Fatal(err)
}

saveResult, err := client.SavePublicResource(
	"https://yadi.sk/d/abc123...",
	stringPtr("saved-document.pdf"),
	stringPtr("/disk/downloads/"),
)
if err != nil {
	log.Fatal(err)
}

func stringPtr(s string) *string {
	return &s
}
```

### Управление корзиной

```go
trash, err := client.GetTrash("/", 50, 0)
if err != nil {
	log.Fatal(err)
}
for _, item := range trash.GetItems() {
	fmt.Printf("- %s (удалён: %s)\n", item.Name, item.Modified)
}

restoreResult, err := client.RestoreFromTrash("/disk/document.pdf", nil, false)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Восстановлено: %s\n", restoreResult.Name)

newName := "restored-document.pdf"
client.RestoreFromTrash("/disk/document.pdf", &newName, false)

err = client.ClearTrash(stringPtr("/disk/document.pdf"))
if err != nil {
	log.Fatal(err)
}

err = client.ClearTrash(nil)
if err != nil {
	log.Fatal(err)
}
```

### Загрузка из URL

```go
uploadResult, err := client.UploadFromURL(
	"https://example.com/document.pdf",
	"/disk/downloads/document.pdf",
	false,
)
if err != nil {
	log.Fatal(err)
}

fmt.Println("Загрузка начата")

if uploadResult.Href != "" {
	operationID := filepath.Base(uploadResult.Href)
	status, err := client.GetOperationStatus(operationID)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Статус: %s\n", status.Status)
	fmt.Printf("Завершено: %v\n", status.IsSuccess())
}
```

### Методы администратора организации

```go
ownedResources, err := client.GetPublicResourcesOwnedByUser(
	"user-uid-123",
	"org-id-456",
	20,
	0,
)
if err != nil {
	log.Fatal(err)
}

for _, resource := range ownedResources.Items {
	fmt.Printf("- %s (%s)\n", resource.Name, resource.PublicURL)
}

accessibleResources, err := client.GetPublicResourcesAccessedByUser(
	"user-uid-123",
	"org-id-456",
	true,
	20,
	nil,
)
if err != nil {
	log.Fatal(err)
}

err = client.UnpublishUserResource(
	"public-key-789",
	"user-uid-123",
	"org-id-456",
)
if err != nil {
	log.Fatal(err)
}
```

## 🔧 Обработка ошибок

SDK предоставляет комплексную обработку ошибок:

```go
result, err := client.UploadFile("/path/to/file.txt", "/disk/file.txt", true)
if err != nil {
	if apiErr, ok := err.(*yandexdisk.APIError); ok {
		fmt.Printf("Ошибка API: %s\n", apiErr.Error())
		fmt.Printf("Код ошибки: %d\n", apiErr.StatusCode)
	} else {
		fmt.Printf("Ошибка: %s\n", err.Error())
	}
	return
}
```

## 📊 Покрытие API

<div align="center">

### 🎯 100% покрытие официального API

</div>

| Категория                | Реализовано | Всего  | Процент  | Статус |
|--------------------------|-------------|--------|----------|--------|
| 💾 Информация о диске    | 1           | 1      | 100%     | ✅      |
| 📁 Операции с файлами    | 8           | 8      | 100%     | ✅      |
| 🔗 Публичные ресурсы     | 8           | 8      | 100%     | ✅      |
| 🗑️ Управление корзиной  | 3           | 3      | 100%     | ✅      |
| 🏷️ Метаданные           | 2           | 2      | 100%     | ✅      |
| 👥 Методы администратора | 3           | 3      | 100%     | ✅      |
| ⚙️ Операции              | 1           | 1      | 100%     | ✅      |
| **🎉 Всего**             | **26**      | **26** | **100%** | **✅**  |

<div align="center">

**Все методы Yandex Disk API реализованы и протестированы!**

</div>

## 🤝 Участие в разработке

<div align="center">

### Мы рады вашему вкладу! 🎉

</div>

Вклады приветствуются! Вот как вы можете помочь:

- 🐛 **Нашли баг?** [Создайте issue](https://github.com/tigusigalpa/yandex-disk-go/issues/new)
- 💡 **Есть идея?** [Предложите улучшение](https://github.com/tigusigalpa/yandex-disk-go/issues/new)
- 🔧 **Хотите помочь?** [Отправьте Pull Request](https://github.com/tigusigalpa/yandex-disk-go/pulls)
- 📖 **Улучшите документацию** - любая помощь ценна!

Прочитайте [CONTRIBUTING.md](CONTRIBUTING.md) для деталей.

## 📄 Лицензия

Этот пакет лицензирован под MIT License. Подробности смотрите в файле [LICENSE](LICENSE).

## 🔗 Полезные ссылки

<table>
<tr>
<td>

### 📚 Документация

- [Yandex Disk API](https://yandex.ru/dev/disk-api/doc/ru/)
- [OAuth Guide](https://yandex.ru/dev/id/doc/ru/)
- [GoDoc](https://pkg.go.dev/github.com/tigusigalpa/yandex-disk-go)

</td>
<td>

### 🛠️ Разработка

- [GitHub Repository](https://github.com/tigusigalpa/yandex-disk-go)
- [Issue Tracker](https://github.com/tigusigalpa/yandex-disk-go/issues)
- [Changelog](https://github.com/tigusigalpa/yandex-disk-go/releases)

</td>
</tr>
</table>

## 📞 Поддержка

Для проблем, вопросов или вкладов:

- Создайте issue на [GitHub](https://github.com/tigusigalpa/yandex-disk-go/issues)
- Проверьте [официальную документацию](https://yandex.ru/dev/disk-api/doc/ru/)
- Ознакомьтесь с [руководством по устранению неполадок](https://yandex.ru/dev/disk-api/doc/ru/concepts/troubleshooting)

---

<div align="center">

### ⭐ Понравился проект?

**Поставьте звезду на GitHub!**

[![GitHub stars](https://img.shields.io/github/stars/tigusigalpa/yandex-disk-go?style=social)](https://github.com/tigusigalpa/yandex-disk-go/stargazers)

---

**Сделано с ❤️ для Go сообщества**

*Автор: [Igor Sazonov](https://github.com/tigusigalpa) | Лицензия: MIT*

</div>
