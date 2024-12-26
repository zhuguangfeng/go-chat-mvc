package dto

type ReviewReq struct {
	ID      int64  `json:"id"`
	Opinion string `json:"opinion"`
	Status  uint   `json:"status"`
}

type ReviewListReq struct {
	BaseListReq
	Biz    string `json:"biz"`   //业务
	BizID  int64  `json:"bizId"` //业务id
	Status uint   `json:"status"`
}

type Review struct {
	ID          int64    `json:"id,omitempty"`
	Biz         string   `json:"biz,omitempty"`
	BizID       int64    `json:"bizId,omitempty"`
	Activity    Activity `json:"activity,omitempty"`
	Reviewer    User     `json:"reviewer,omitempty"`
	Status      uint     `json:"status,omitempty"`
	CreatedTime string   `json:"createdTime,omitempty"`
	ReviewTime  string   `json:"reviewTime,omitempty"`
}
