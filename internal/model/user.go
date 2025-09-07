package model

type User struct {
	Username  string `json:"username"`
	SessionID string `json:"session_id"`
}
