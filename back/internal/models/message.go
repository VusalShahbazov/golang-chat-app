package models

type Message struct {
	Model
	Body string `json:"body"`
	FromUserId uint64 `json:"from_user_id"`
	FromUser User `json:"from_user" gorm:"foreignKey:FromUserId"`
	RoomId uint64 `json:"room_id"`
	Room *Room `json:"room,omitempty" gorm:"foreignKey:FromUserId"`
}