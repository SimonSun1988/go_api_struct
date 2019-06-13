package models

type Role struct {
    Id int `json:"id" gorm:"column:Id; primary_key"`
    Name string `json:"name" gorm:"column:Name; type:varchar(100)"`
}