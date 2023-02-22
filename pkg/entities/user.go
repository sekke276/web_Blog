package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json: "id" bson : "_id, omitempty"`
	Username  string             `json:"userName" bson: "username, omitempty"`
	Password  []byte             `json:"password" bson: "password, omitempty"`
	Birthdate time.Time          `json:"bdate" bson: "bdate, omitempty"`
	Avatar    string             `json:"avt" bson: "avatar`
	Facebook  string             `json: "facebook" bson: "facebook"`
	Gmail     string
	Gender    string `json:"gender" bson: "gender"`
}
