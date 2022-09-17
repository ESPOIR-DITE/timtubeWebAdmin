package domain

type Video struct {
	Id           string  `json:"id" gorm:"primaryKey"`
	Title        string  `json:"title"`
	Date         string  `json:"date"`
	DateUploaded string  `json:"dateUploaded"`
	Description  string  `json:"description"`
	IsPrivate    bool    `json:"isPrivate"`
	Price        float64 `json:"price"`
	URL          string  `json:"url"`
}
type VideoThriller struct {
	Id          string `json:"id"`
	VideoId     string `json:"video_id"`
	Description string `json:"description"`
}
type VideoData struct {
	Id       string `json:"id" gorm:"primaryKey"`
	Picture  []byte `json:"picture"`
	Video    []byte `json:"video"`
	FileType string `json:"fileType"`
	FileSize string `json:"fileSize"`
}

type VideoCategory struct {
	Id         string `json:"id" gorm:"primaryKey"`
	VideoId    string `json:"videoId"`
	CategoryId string `json:"categoryId"`
}

type VideoComment struct {
	Id        string `json:"id" gorm:"primaryKey"`
	VideoId   string `json:"videoId"`
	CommentId string `json:"categoryId"`
}

type Category struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type VideoRelated struct {
	Id             string `json:"id" gorm:"primaryKey"`
	CurrentVideoId string `json:"currentVideo"`
	RelatedVideoId string `json:"relatedVideoId"`
}
