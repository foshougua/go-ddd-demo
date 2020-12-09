package entity

type Article struct {
	Author     string `json:"author"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	UpTime     int64  `json:"up_time"`
	DelTime    int64  `json:"del_time"`
}
