package models

type Ratio struct {
	UserID      string `bson:"userid" json:"userId"`
	UserRatioID string `bson:"userratioid" json:"userRatioId"`
}
