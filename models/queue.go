package models

type Queue struct {
	Id        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	GoToQueue string `json:"go_to_queue"`
}
