package entities

import (
	"time"
)

type User struct {
	Id        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string    `json:"userName" bson: "username, omitempty"`
	Password  string    `json:"password" bson: "password, omitempty"`
	Birthdate time.Time `json:"bdate" bson: "bdate, omitempty"`
	Avatar    string    `json:"avt" bson: "avatar`
	Facebook  string    `json: "facebook" bson: "facebook"`
	Gender    string    `json:"gender" bson: "gender"`
}
type UserRequest struct {
	Username  string    `json:"username" validate: "required,min=8,max=32"`
	Password  string    `json:"password" validate: "required, min=8, max=32"`
	Gender    string    `json: "gender" validate: "required, oneof= Male Female"`
	Birthdate time.Time `json:"bdate"`
	Avatar    string    `json:"avatar" validate: "required, datauri"`
	Facebook  string    `json: "facebook" validate:"datauri"`
}
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password`
}
