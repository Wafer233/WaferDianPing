package application

type VoucherVO struct {
	Id          int64  `json:"id"`
	ShopId      int64  `json:"shopId"`
	Title       string `json:"title"`
	SubTitle    string `json:"subTitle"`
	Rules       string `json:"rules"`
	PayValue    int64  `json:"payValue"`
	ActualValue int64  `json:"actualValue"`
	Type        int    `json:"type"`
	Status      int    `json:"status"`
	Stock       int    `json:"stock"`
	BeginTime   string `json:"beginTime"`
	EndTime     string `json:"endTime"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}
