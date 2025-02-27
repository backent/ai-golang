package models

import "time"

var QuestionTable string = "questions"

type Question struct {
	Id            int64
	Name          string
	Username      string
	Amount        int
	GeminiFileURI string
	FileName      string
	Result        string
	Exams         []Exam
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
