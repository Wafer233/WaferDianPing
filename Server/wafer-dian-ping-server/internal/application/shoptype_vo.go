package application

type ShopTypeVO struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Sort       int    `json:"sort"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
