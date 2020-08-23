package models

import "time"

type Message struct {
	Username string
	UserId string
	MessageText string
	IsSub bool
	IsMod bool
	IsTurbo bool
	Channel string
	DateSent time.Time
}