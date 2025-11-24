package application

type ShopVO struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	TypeId     int64   `json:"typeId"`
	Images     string  `json:"images"`
	Area       string  `json:"area"`
	Address    string  `json:"address"`
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	AvgPrice   int64   `json:"avgPrice"`
	Sold       int     `json:"sold"`
	Comments   int     `json:"comments"`
	Score      int     `json:"score"`
	OpenHours  string  `json:"openHours"`
	CreateTime string  `json:"createTime"`
	UpdateTime string  `json:"updateTime"`
	Distance   int64   `json:"distance"`
}
