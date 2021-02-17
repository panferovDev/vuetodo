package models

type Task struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Titile string `json:"title"`
	Done  bool `json:"done"`
}