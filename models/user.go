package models

type User struct {
	Id       uint   //'json:"id"'
	UserID   string //'json:"user_id gorm:"unique"'
	Password []byte //json:"-"
}
