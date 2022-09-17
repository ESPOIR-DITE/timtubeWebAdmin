package domain

type Dashboard struct {
	PendingUsers int64 `json:"pending_users"`
	Users        int64 `json:"users"`
	Videos       int64 `json:"videos"`
	Channels     int64 `json:"channels"`
	UserStack    UserStack
}
