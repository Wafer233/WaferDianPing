package application

type UserVO struct {
	Id       int64  `json:"id"`
	NickName string `json:"nickName"`
	Icon     string `json:"icon"`
}

type UserInfoVO struct {
	UserId     int64  `json:"userId"`
	City       string `json:"city"`
	Introduce  string `json:"introduce"`
	Fans       int    `json:"fans"`
	Followee   int    `json:"followee"`
	Gender     bool   `json:"gender"`
	Birthday   string `json:"birthday"`
	Credits    int    `json:"credits"`
	Level      int    `json:"level" `
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
