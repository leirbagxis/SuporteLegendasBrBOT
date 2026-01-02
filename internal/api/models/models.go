package models

type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Version string `json:"version"`
}

type BotInfo struct {
	Username string `json:"username"`
	ID       int64  `json:"id"`
	IsActive bool   `json:"is_active"`
}

type CaptionRequest struct {
	ChannelID int64  `json:"channel_id"`
	MessageID int    `json:"message_id"`
	Caption   string `json:"caption"`
}

type CaptionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type InvitationRequest struct {
	InvitationLink string `json:"invitation_link"`
}
