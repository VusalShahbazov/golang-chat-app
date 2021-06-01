package models


type Room struct {
	Model
	Name string `json:"name"`
	OwnerId uint64 `json:"owner_id"`
 	Owner User `json:"owner" gorm:"foreignKey:OwnerId"`
}