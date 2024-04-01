package entity

// 名人说的话
type Say struct {
	Id         string `json:"id"`
	PersonName string `json:"personName"`
	Content    string `json:"content"`
	Avatar     string `json:"avatar"`
	Image      string `json:"image"`
	LikeNum    string `json:"likeNum"`
	CommentNum string `json:"commentNum"`
	RecordTime string `json:"recordTime"`
}

type Comment struct {
	SayId      string `json:"sayId"`
	NickName   string `json:"nickName"`
	RecordTime string `json:"recordTime"`
	Content    string `json:"content"`
}
