package context

import "path"

// 定义响应状态数据
const (
	StatusOK    = "OK"
	StatusError = "error"
	StatusFail  = "fail"
)

// MRouterTitle 路由关联的标题数据
var MRouterTitle = make(map[string]string)

// GetRouterTitleKey 获取路由标题键名
func GetRouterTitleKey(method, router string) string {
	return path.Join(method, router)
}

// HTTPError HTTP响应错误
type HTTPError struct {
	Error HTTPErrorItem `json:"error,omitempty"`
}

// HTTPErrorItem HTTP响应错误项
type HTTPErrorItem struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// HTTPStatus HTTP响应状态
type HTTPStatus struct {
	Status string `json:"status,omitempty"`
}

// HTTPList HTTP响应列表数据
type HTTPList struct {
	List       interface{}     `json:"list,omitempty"`
	Pagination *HTTPPagination `json:"pagination,omitempty"`
}

// HTTPPagination HTTP分页数据
type HTTPPagination struct {
	Total    int64 `json:"total,omitempty"`
	Current  uint  `json:"current,omitempty"`
	PageSize uint  `json:"pageSize,omitempty"`
}
