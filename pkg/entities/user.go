package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `json: "id" bson : "_id, omitempty"`
	username  string             `json:"userName" bson: "username, omitempty"`
	password  []byte             `json:"password" bson: "password, omitempty"`
	birthdate time.Time          `json:"bdate" bson: "bdate, omitempty"`
	avatar    string             `json:"avt" bson: "avatar`
	facebook  string             `json: "facebook" bson: "facebook"`
	gmail     string
	gender    string `json:"gender" bson: "gender"`
}
