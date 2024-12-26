package dto

type BaseDeleteReq struct {
	ID int64 `json:"id" dc:"活动id"`
}

type BaseListReq struct {
	PageNum   int      `json:"pageNum"`
	PageSize  int      `json:"pageSize"`
	SearchKey string   `json:"searchKey"`
	Order     []string `json:"order"`
}
