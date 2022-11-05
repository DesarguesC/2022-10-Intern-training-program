package main
// table
type Users struct {
	//gorm.Model
	Id uint `gorm:"primary_key"`
	Name string
	Passwd string
}
type Todos struct{
	Id uint `gorm:"promary_key"`
	User_id uint
	tittle string
	content string
}