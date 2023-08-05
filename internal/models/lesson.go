package models

type Lesson struct {
	ID             int      `json:"id"`
	Title          string   `json:"title"`
	BlockingLesson int      `json:"blocking_lesson"`
	BlockingText   string   `json:"blocking_text"`
	Tags           []string `json:"tags"`
}
