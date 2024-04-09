package models

const (
	TableUser string = "user"
)

func (u *User) TableName() string { return TableUser }
