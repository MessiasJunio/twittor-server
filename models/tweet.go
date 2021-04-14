package models

//Tweet capture the body, the message that come to us
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
