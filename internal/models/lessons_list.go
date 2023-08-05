package models

type ListLesson struct {
	Lesson
	IsCompleted bool `json:"is_completed"`
	IsOpen      bool `json:"is_open"`
}
