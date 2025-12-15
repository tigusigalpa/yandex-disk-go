package yandexdisk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationURL(t *testing.T) {
	clientID := "test-client-id"
	url := GetAuthorizationURL(clientID)

	assert.Contains(t, url, "https://oauth.yandex.ru/authorize")
	assert.Contains(t, url, "response_type=token")
	assert.Contains(t, url, "client_id=test-client-id")
}

func TestNewClient(t *testing.T) {
	token := "test-token"
	client := NewClient(token)

	assert.NotNil(t, client)
	assert.Equal(t, token, client.accessToken)
	assert.NotNil(t, client.httpClient)
}

func TestDiskInfoGetFreeSpace(t *testing.T) {
	diskInfo := &DiskInfo{
		TotalSpace: 10000000000,
		UsedSpace:  3000000000,
	}

	freeSpace := diskInfo.GetFreeSpace()
	assert.Equal(t, int64(7000000000), freeSpace)
}

func TestDiskInfoGetUsagePercentage(t *testing.T) {
	diskInfo := &DiskInfo{
		TotalSpace: 10000000000,
		UsedSpace:  3000000000,
	}

	percentage := diskInfo.GetUsagePercentage()
	assert.Equal(t, 30.0, percentage)
}

func TestDiskInfoGetUsagePercentageZeroTotal(t *testing.T) {
	diskInfo := &DiskInfo{
		TotalSpace: 0,
		UsedSpace:  0,
	}

	percentage := diskInfo.GetUsagePercentage()
	assert.Equal(t, 0.0, percentage)
}

func TestResourceIsDir(t *testing.T) {
	resource := &Resource{Type: "dir"}
	assert.True(t, resource.IsDir())
	assert.False(t, resource.IsFile())
}

func TestResourceIsFile(t *testing.T) {
	resource := &Resource{Type: "file"}
	assert.True(t, resource.IsFile())
	assert.False(t, resource.IsDir())
}

func TestResourceIsPublished(t *testing.T) {
	resource := &Resource{PublicURL: "https://yadi.sk/d/abc123"}
	assert.True(t, resource.IsPublished())

	resource2 := &Resource{PublicKey: "key123"}
	assert.True(t, resource2.IsPublished())

	resource3 := &Resource{}
	assert.False(t, resource3.IsPublished())
}

func TestResourceGetItems(t *testing.T) {
	resource := &Resource{
		Embedded: &Embedded{
			Items: []Resource{
				{Name: "file1.txt"},
				{Name: "file2.txt"},
			},
		},
	}

	items := resource.GetItems()
	assert.Len(t, items, 2)
	assert.Equal(t, "file1.txt", items[0].Name)
	assert.Equal(t, "file2.txt", items[1].Name)
}

func TestResourceGetItemsEmpty(t *testing.T) {
	resource := &Resource{}
	items := resource.GetItems()
	assert.Len(t, items, 0)
}

func TestResourceGetTotalItems(t *testing.T) {
	resource := &Resource{
		Embedded: &Embedded{
			Total: 42,
		},
	}

	total := resource.GetTotalItems()
	assert.Equal(t, 42, total)
}

func TestResourceGetTotalItemsEmpty(t *testing.T) {
	resource := &Resource{}
	total := resource.GetTotalItems()
	assert.Equal(t, 0, total)
}

func TestOperationIsInProgress(t *testing.T) {
	op := &Operation{Status: "in-progress"}
	assert.True(t, op.IsInProgress())
	assert.False(t, op.IsSuccess())
	assert.False(t, op.IsFailed())
}

func TestOperationIsSuccess(t *testing.T) {
	op := &Operation{Status: "success"}
	assert.False(t, op.IsInProgress())
	assert.True(t, op.IsSuccess())
	assert.False(t, op.IsFailed())
}

func TestOperationIsFailed(t *testing.T) {
	op := &Operation{Status: "failed"}
	assert.False(t, op.IsInProgress())
	assert.False(t, op.IsSuccess())
	assert.True(t, op.IsFailed())
}

func TestAPIErrorError(t *testing.T) {
	err := &APIError{
		Message:     "Test message",
		Description: "Test description",
		StatusCode:  400,
	}

	assert.Equal(t, "Test description", err.Error())

	err2 := &APIError{
		Message:    "Test message",
		StatusCode: 400,
	}
	assert.Equal(t, "Test message", err2.Error())

	err3 := &APIError{
		Error:      "Test error",
		StatusCode: 400,
	}
	assert.Equal(t, "Test error", err3.Error())

	err4 := &APIError{
		StatusCode: 400,
	}
	assert.Equal(t, "unknown API error", err4.Error())
}
