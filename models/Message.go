package models

type Message struct {
	Username string
	UserId string
	MessageText string
	IsSub bool
	IsMod bool
	IsTurbo bool
}