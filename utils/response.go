package util

type ApiResponse struct {
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

type PaginatioResponse struct {
	Data interface{}  `json:"data,omitempty"`
	Meta PageMetadata `json:"meta,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"totalItem"`
	TotalPage int64 `json:"totalPage"`
}
