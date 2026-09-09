package models

type Schedule struct {
	Id               int    `json:"id"`
	Day              string `json:"day"`
	EventDescription string `json:"eventDescription"`
}
