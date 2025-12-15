package yandexdisk

type DiskInfo struct {
	TotalSpace                  int64             `json:"total_space"`
	UsedSpace                   int64             `json:"used_space"`
	TrashSize                   int64             `json:"trash_size"`
	UnlimitedAutouploadEnabled  bool              `json:"unlimited_autoupload_enabled"`
	MaxFileSize                 int64             `json:"max_file_size"`
	PaidMaxFileSize             int64             `json:"paid_max_file_size"`
	IsPaid                      bool              `json:"is_paid"`
	SystemFolders               map[string]string `json:"system_folders"`
	User                        User              `json:"user"`
	Revision                    int64             `json:"revision"`
}

func (d *DiskInfo) GetFreeSpace() int64 {
	return d.TotalSpace - d.UsedSpace
}

func (d *DiskInfo) GetUsagePercentage() float64 {
	if d.TotalSpace == 0 {
		return 0.0
	}
	return float64(d.UsedSpace) / float64(d.TotalSpace) * 100
}

type User struct {
	Country     string `json:"country"`
	Login       string `json:"login"`
	DisplayName string `json:"display_name"`
	UID         string `json:"uid"`
}

type Resource struct {
	Name             string                 `json:"name"`
	Path             string                 `json:"path"`
	Type             string                 `json:"type"`
	Size             int64                  `json:"size"`
	Created          string                 `json:"created"`
	Modified         string                 `json:"modified"`
	MimeType         string                 `json:"mime_type"`
	MediaType        string                 `json:"media_type"`
	Preview          string                 `json:"preview"`
	File             string                 `json:"file"`
	MD5              string                 `json:"md5"`
	SHA256           string                 `json:"sha256"`
	PublicURL        string                 `json:"public_url"`
	PublicKey        string                 `json:"public_key"`
	ResourceID       string                 `json:"resource_id"`
	CustomProperties map[string]interface{} `json:"custom_properties"`
	CommentIDs       CommentIDs             `json:"comment_ids"`
	Owner            Owner                  `json:"owner"`
	Embedded         *Embedded              `json:"_embedded,omitempty"`
	Revision         int64                  `json:"revision"`
}

func (r *Resource) IsDir() bool {
	return r.Type == "dir"
}

func (r *Resource) IsFile() bool {
	return r.Type == "file"
}

func (r *Resource) IsPublished() bool {
	return r.PublicURL != "" || r.PublicKey != ""
}

func (r *Resource) GetItems() []Resource {
	if r.Embedded == nil {
		return []Resource{}
	}
	return r.Embedded.Items
}

func (r *Resource) GetTotalItems() int {
	if r.Embedded == nil {
		return 0
	}
	return r.Embedded.Total
}

type CommentIDs struct {
	PrivateResource string `json:"private_resource"`
	PublicResource  string `json:"public_resource"`
}

type Owner struct {
	Login       string `json:"login"`
	DisplayName string `json:"display_name"`
	UID         string `json:"uid"`
}

type Embedded struct {
	Sort   string     `json:"sort"`
	Path   string     `json:"path"`
	Items  []Resource `json:"items"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
	Total  int        `json:"total"`
}

type FilesList struct {
	Items  []Resource `json:"items"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
	Total  int        `json:"total"`
}

type UploadResult struct {
	Status  int
	Success bool
}

type Operation struct {
	Status    string `json:"status"`
	Type      string `json:"type"`
	Href      string `json:"href"`
	Method    string `json:"method"`
	Templated bool   `json:"templated"`
}

func (o *Operation) IsInProgress() bool {
	return o.Status == "in-progress"
}

func (o *Operation) IsSuccess() bool {
	return o.Status == "success"
}

func (o *Operation) IsFailed() bool {
	return o.Status == "failed"
}

type APIError struct {
	Message     string `json:"message"`
	Description string `json:"description"`
	Error       string `json:"error"`
	StatusCode  int
}

func (e *APIError) Error() string {
	if e.Description != "" {
		return e.Description
	}
	if e.Message != "" {
		return e.Message
	}
	if e.Error != "" {
		return e.Error
	}
	return "unknown API error"
}
