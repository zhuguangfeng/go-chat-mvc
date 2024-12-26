package dto

type CreateUserReq struct {
	//Username        string `json:"username"`
	//Avatar          string `json:"avatar"`
	//BackgroundImage string `json:"backgroundImage"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type PasswordLoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type User struct {
	ID              int64  `json:"id,omitempty"`
	Username        string `json:"username,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	BackgroundImage string `json:"backgroundImage,omitempty"`
	Gender          uint   `json:"gender,omitempty"`
	Age             uint   `json:"age,omitempty"`
	Status          uint   `json:"status,omitempty"`
	CreatedTime     string `json:"createdTime,omitempty"`
}
