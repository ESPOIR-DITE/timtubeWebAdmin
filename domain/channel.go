package domain

type Channel struct {
	Id            string `json:"id" `
	Name          string `json:"name"`
	ChannelTypeId string `json:"channel_type_id"`
	UserId        string `json:"user_id"`
	Region        string `json:"region"`
	Date          string `json:"date"`
	Image         []byte `json:"image"`
	ImageBase64   string `json:"image_base_64"`
	Description   string `json:"description"`
}
type ChannelVideos struct {
	Id          string `json:"id"`
	VideoId     string `json:"video_id"`
	ChannelId   string `json:"channel_id"`
	Description string `json:"description"`
}

type ChannelType struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ChannelSubscription struct {
	Id        string `json:"id"`
	ChannelId string `json:"channel_id"`
	UserId    string `json:"user_id"`
	Date      string `json:"date"`
}
