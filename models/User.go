package models

type User struct {
	Id int `form:"id"`
	Phone string `form:"phone"`
	Pwd string `form:"pwd"`
}
