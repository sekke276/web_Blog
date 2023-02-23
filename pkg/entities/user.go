package entities

import (
	"time"
)

type User struct {
	Username  string    `json:"userName" bson: "username, omitempty"`
	Password  string    `json:"password" bson: "password, omitempty"`
	Birthdate time.Time `json:"bdate" bson: "bdate, omitempty"`
	Avatar    string    `json:"avt" bson: "avatar`
	Facebook  string    `json: "facebook" bson: "facebook"`
	Gender    string    `json:"gender" bson: "gender"`
}
