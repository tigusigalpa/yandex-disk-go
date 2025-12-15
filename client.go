package yandexdisk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

const (
	APIBaseURL = "https://cloud-api.yandex.net/v1/disk"
)

type Client struct {
	accessToken string
	httpClient  *http.Client
}

func NewClient(accessToken string) *Client {
	return &Client{
		accessToken: accessToken,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

func GetAuthorizationURL(clientID string) string {
	params := url.Values{}
	params.Set("response_type", "token")
	params.Set("client_id", clientID)
	return "https://oauth.yandex.ru/authorize?" + params.Encode()
}

func (c *Client) request(method, endpoint string, queryParams url.Values, body interface{}) ([]byte, error) {
	reqURL := APIBaseURL + endpoint
	if queryParams != nil {
		reqURL += "?" + queryParams.Encode()
	}

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, reqURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "OAuth "+c.accessToken)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var apiError APIError
		if err := json.Unmarshal(respBody, &apiError); err == nil && apiError.Message != "" {
			apiError.StatusCode = resp.StatusCode
			return nil, &apiError
		}
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func (c *Client) GetCapacity() (*DiskInfo, error) {
	data, err := c.request("GET", "/", nil, nil)
	if err != nil {
		return nil, err
	}

	var diskInfo DiskInfo
	if err := json.Unmarshal(data, &diskInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal disk info: %w", err)
	}

	return &diskInfo, nil
}

func (c *Client) GetMeta(path string, params map[string]string) (*Resource, error) {
	queryParams := url.Values{}
	queryParams.Set("path", path)
	for k, v := range params {
		queryParams.Set(k, v)
	}

	data, err := c.request("GET", "/resources", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var resource Resource
	if err := json.Unmarshal(data, &resource); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource: %w", err)
	}

	return &resource, nil
}

func (c *Client) AddMeta(path string, customProperties map[string]interface{}) (*Resource, error) {
	body := map[string]interface{}{
		"path":              path,
		"custom_properties": customProperties,
	}

	data, err := c.request("PATCH", "/resources", nil, body)
	if err != nil {
		return nil, err
	}

	var resource Resource
	if err := json.Unmarshal(data, &resource); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource: %w", err)
	}

	return &resource, nil
}

func (c *Client) GetAllFiles(limit, offset int) (*FilesList, error) {
	queryParams := url.Values{}
	queryParams.Set("limit", fmt.Sprintf("%d", limit))
	queryParams.Set("offset", fmt.Sprintf("%d", offset))

	data, err := c.request("GET", "/resources/files", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var filesList FilesList
	if err := json.Unmarshal(data, &filesList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal files list: %w", err)
	}

	return &filesList, nil
}

func (c *Client) GetRecentUploads(limit, offset int) (*FilesList, error) {
	queryParams := url.Values{}
	queryParams.Set("limit", fmt.Sprintf("%d", limit))
	queryParams.Set("offset", fmt.Sprintf("%d", offset))

	data, err := c.request("GET", "/resources/last-uploaded", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var filesList FilesList
	if err := json.Unmarshal(data, &filesList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal files list: %w", err)
	}

	return &filesList, nil
}

func (c *Client) GetRecentPublished(limit, offset int) (*FilesList, error) {
	queryParams := url.Values{}
	queryParams.Set("limit", fmt.Sprintf("%d", limit))
	queryParams.Set("offset", fmt.Sprintf("%d", offset))

	data, err := c.request("GET", "/resources/public", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var filesList FilesList
	if err := json.Unmarshal(data, &filesList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal files list: %w", err)
	}

	return &filesList, nil
}

func (c *Client) CreateFolder(path string) (*Resource, error) {
	queryParams := url.Values{}
	queryParams.Set("path", path)

	data, err := c.request("PUT", "/resources", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var resource Resource
	if err := json.Unmarshal(data, &resource); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource: %w", err)
	}

	return &resource, nil
}

func (c *Client) UploadFile(localFilePath, remotePath string, overwrite bool) (*UploadResult, error) {
	if _, err := os.Stat(localFilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("local file not found: %s", localFilePath)
	}

	queryParams := url.Values{}
	queryParams.Set("path", remotePath)
	if overwrite {
		queryParams.Set("overwrite", "true")
	} else {
		queryParams.Set("overwrite", "false")
	}

	data, err := c.request("GET", "/resources/upload", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var uploadURL struct {
		Href      string `json:"href"`
		Method    string `json:"method"`
		Templated bool   `json:"templated"`
	}
	if err := json.Unmarshal(data, &uploadURL); err != nil {
		return nil, fmt.Errorf("failed to unmarshal upload URL: %w", err)
	}

	if uploadURL.Href == "" {
		return nil, fmt.Errorf("failed to get upload URL")
	}

	fileContent, err := os.ReadFile(localFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read local file: %w", err)
	}

	req, err := http.NewRequest("PUT", uploadURL.Href, bytes.NewBuffer(fileContent))
	if err != nil {
		return nil, fmt.Errorf("failed to create upload request: %w", err)
	}

	req.Header.Set("Authorization", "OAuth "+c.accessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("upload failed: %w", err)
	}
	defer resp.Body.Close()

	return &UploadResult{
		Status:  resp.StatusCode,
		Success: resp.StatusCode == 201,
	}, nil
}

func (c *Client) DownloadFile(remotePath, localPath string) error {
	queryParams := url.Values{}
	queryParams.Set("path", remotePath)

	data, err := c.request("GET", "/resources/download", queryParams, nil)
	if err != nil {
		return err
	}

	var downloadURL struct {
		Href   string `json:"href"`
		Method string `json:"method"`
	}
	if err := json.Unmarshal(data, &downloadURL); err != nil {
		return fmt.Errorf("failed to unmarshal download URL: %w", err)
	}

	if downloadURL.Href == "" {
		return fmt.Errorf("failed to get download URL for: %s", remotePath)
	}

	req, err := http.NewRequest("GET", downloadURL.Href, nil)
	if err != nil {
		return fmt.Errorf("failed to create download request: %w", err)
	}

	req.Header.Set("Authorization", "OAuth "+c.accessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("download failed with status: %d", resp.StatusCode)
	}

	if err := os.MkdirAll(filepath.Dir(localPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	outFile, err := os.Create(localPath)
	if err != nil {
		return fmt.Errorf("failed to create local file: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (c *Client) Copy(fromPath, toPath string, overwrite bool) (*Resource, error) {
	queryParams := url.Values{}
	queryParams.Set("from", fromPath)
	queryParams.Set("path", toPath)
	if overwrite {
		queryParams.Set("overwrite", "true")
	} else {
		queryParams.Set("overwrite", "false")
	}

	data, err := c.request("POST", "/resources/copy", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var resource Resource
	if err := json.Unmarshal(data, &resource); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource: %w", err)
	}

	return &resource, nil
}

func (c *Client) Move(fromPath, toPath string, overwrite bool) (*Resource, error) {
	queryParams := url.Values{}
	queryParams.Set("from", fromPath)
	queryParams.Set("path", toPath)
	if overwrite {
		queryParams.Set("overwrite", "true")
	} else {
		queryParams.Set("overwrite", "false")
	}

	data, err := c.request("POST", "/resources/move", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var resource Resource
	if err := json.Unmarshal(data, &resource); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource: %w", err)
	}

	return &resource, nil
}

func (c *Client) Delete(path string, permanently bool) error {
	queryParams := url.Values{}
	queryParams.Set("path", path)
	if permanently {
		queryParams.Set("permanently", "true")
	} else {
		queryParams.Set("permanently", "false")
	}

	_, err := c.request("DELETE", "/resources", queryParams, nil)
	return err
}

func (c *Client) Publish(path string) (*Resource, error) {
	queryParams := url.Values{}
	queryParams.Set("path", path)

	data, err := c.request("PUT", "/resources/publish", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var resource Resource
	if err := json.Unmarshal(data, &resource); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource: %w", err)
	}

	return &resource, nil
}

func (c *Client) Unpublish(path string) (*Resource, error) {
	queryParams := url.Values{}
	queryParams.Set("path", path)

	data, err := c.request("PUT", "/resources/unpublish", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var resource Resource
	if err := json.Unmarshal(data, &resource); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource: %w", err)
	}

	return &resource, nil
}

func (c *Client) GetPublicResourceMeta(publicKey string, params map[string]string) (*Resource, error) {
	queryParams := url.Values{}
	queryParams.Set("public_key", publicKey)
	for k, v := range params {
		queryParams.Set(k, v)
	}

	data, err := c.request("GET", "/public/resources", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var resource Resource
	if err := json.Unmarshal(data, &resource); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource: %w", err)
	}

	return &resource, nil
}

func (c *Client) DownloadPublicResource(publicKey, localPath string, path *string) error {
	queryParams := url.Values{}
	queryParams.Set("public_key", publicKey)
	if path != nil {
		queryParams.Set("path", *path)
	}

	data, err := c.request("GET", "/public/resources/download", queryParams, nil)
	if err != nil {
		return err
	}

	var downloadURL struct {
		Href   string `json:"href"`
		Method string `json:"method"`
	}
	if err := json.Unmarshal(data, &downloadURL); err != nil {
		return fmt.Errorf("failed to unmarshal download URL: %w", err)
	}

	if downloadURL.Href == "" {
		return fmt.Errorf("failed to get download URL for public resource")
	}

	resp, err := c.httpClient.Get(downloadURL.Href)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("download failed with status: %d", resp.StatusCode)
	}

	if err := os.MkdirAll(filepath.Dir(localPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	outFile, err := os.Create(localPath)
	if err != nil {
		return fmt.Errorf("failed to create local file: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (c *Client) SavePublicResource(publicKey string, name, path *string) (*Resource, error) {
	queryParams := url.Values{}
	queryParams.Set("public_key", publicKey)
	if name != nil {
		queryParams.Set("name", *name)
	}
	if path != nil {
		queryParams.Set("path", *path)
	}

	data, err := c.request("POST", "/public/resources/save", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var resource Resource
	if err := json.Unmarshal(data, &resource); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource: %w", err)
	}

	return &resource, nil
}

func (c *Client) GetAvailablePublicSettings() (map[string]interface{}, error) {
	data, err := c.request("GET", "/public/resources/public-settings/available", nil, nil)
	if err != nil {
		return nil, err
	}

	var settings map[string]interface{}
	if err := json.Unmarshal(data, &settings); err != nil {
		return nil, fmt.Errorf("failed to unmarshal settings: %w", err)
	}

	return settings, nil
}

func (c *Client) GetPublicSettings(path string, allowAddressAccess bool) (map[string]interface{}, error) {
	queryParams := url.Values{}
	queryParams.Set("path", path)
	if allowAddressAccess {
		queryParams.Set("allow_address_access", "true")
	}

	data, err := c.request("GET", "/public/resources/public-settings", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var settings map[string]interface{}
	if err := json.Unmarshal(data, &settings); err != nil {
		return nil, fmt.Errorf("failed to unmarshal settings: %w", err)
	}

	return settings, nil
}

func (c *Client) ChangePublicSettings(path string, settings map[string]interface{}) (map[string]interface{}, error) {
	queryParams := url.Values{}
	queryParams.Set("path", path)

	data, err := c.request("PUT", "/resources/public", queryParams, settings)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal result: %w", err)
	}

	return result, nil
}

func (c *Client) UploadFromURL(fileURL, remotePath string, disableRedirects bool) (*Operation, error) {
	queryParams := url.Values{}
	queryParams.Set("url", fileURL)
	queryParams.Set("path", remotePath)
	if disableRedirects {
		queryParams.Set("disable_redirects", "true")
	}

	data, err := c.request("POST", "/resources/upload", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var operation Operation
	if err := json.Unmarshal(data, &operation); err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation: %w", err)
	}

	return &operation, nil
}

func (c *Client) GetOperationStatus(operationID string) (*Operation, error) {
	data, err := c.request("GET", "/operations/"+operationID, nil, nil)
	if err != nil {
		return nil, err
	}

	var operation Operation
	if err := json.Unmarshal(data, &operation); err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation: %w", err)
	}

	return &operation, nil
}

func (c *Client) GetTrash(path string, limit, offset int) (*Resource, error) {
	queryParams := url.Values{}
	queryParams.Set("path", path)
	queryParams.Set("limit", fmt.Sprintf("%d", limit))
	queryParams.Set("offset", fmt.Sprintf("%d", offset))

	data, err := c.request("GET", "/trash/resources", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var resource Resource
	if err := json.Unmarshal(data, &resource); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource: %w", err)
	}

	return &resource, nil
}

func (c *Client) RestoreFromTrash(path string, name *string, overwrite bool) (*Resource, error) {
	queryParams := url.Values{}
	queryParams.Set("path", path)
	if name != nil {
		queryParams.Set("name", *name)
	}
	if overwrite {
		queryParams.Set("overwrite", "true")
	}

	data, err := c.request("PUT", "/trash/resources/restore", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var resource Resource
	if err := json.Unmarshal(data, &resource); err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource: %w", err)
	}

	return &resource, nil
}

func (c *Client) ClearTrash(path *string) error {
	queryParams := url.Values{}
	if path != nil {
		queryParams.Set("path", *path)
	}

	_, err := c.request("DELETE", "/trash/resources", queryParams, nil)
	return err
}

func (c *Client) GetPublicResourcesOwnedByUser(userID, orgID string, limit, offset int) (*FilesList, error) {
	queryParams := url.Values{}
	queryParams.Set("user_id", userID)
	queryParams.Set("org_id", orgID)
	queryParams.Set("limit", fmt.Sprintf("%d", limit))
	queryParams.Set("offset", fmt.Sprintf("%d", offset))

	data, err := c.request("GET", "/public/resources/admin/public-resources", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var filesList FilesList
	if err := json.Unmarshal(data, &filesList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal files list: %w", err)
	}

	return &filesList, nil
}

func (c *Client) GetPublicResourcesAccessedByUser(userID, orgID string, includeGroupAccess bool, limit int, iterationKey *string) (*FilesList, error) {
	queryParams := url.Values{}
	queryParams.Set("user_id", userID)
	queryParams.Set("org_id", orgID)
	queryParams.Set("limit", fmt.Sprintf("%d", limit))
	if includeGroupAccess {
		queryParams.Set("include_group_access", "true")
	}
	if iterationKey != nil {
		queryParams.Set("iteration_key", *iterationKey)
	}

	data, err := c.request("GET", "/public/resources/admin/accessible-resources", queryParams, nil)
	if err != nil {
		return nil, err
	}

	var filesList FilesList
	if err := json.Unmarshal(data, &filesList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal files list: %w", err)
	}

	return &filesList, nil
}

func (c *Client) UnpublishUserResource(publicKey, userID, orgID string) error {
	queryParams := url.Values{}
	queryParams.Set("public_key", publicKey)
	queryParams.Set("user_id", userID)
	queryParams.Set("org_id", orgID)

	_, err := c.request("PUT", "/public/resources/admin/unpublish", queryParams, nil)
	return err
}
