package v1

type MetricsPluginsReq struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	PageSize int32  `json:"page_size"`
	PageNum  int32  `json:"page_num"`
}
