package application

type BlogDTO struct {
	Id         int64  `json:"id"`
	ShopId     int64  `json:"shopId"`
	UserId     int64  `json:"userId"`
	Icon       string `json:"icon"`
	Name       string `json:"name"`
	IsLike     bool   `json:"isLike"`
	Title      string `json:"title"`
	Images     string `json:"images"`
	Content    string `json:"content"`
	Liked      int    `json:"liked"`
	Comments   int    `json:"comments"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
