package models

import (
	"database/sql"
	"time"
)

var QuestionTable string = "questions"

type Question struct {
	Id            int64
	Name          string
	Username      string
	Chapter       sql.NullString
	Amount        int
	GeminiFileURI string
	FileName      string
	Result        string
	Exams         []Exam
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type NullAbleExam struct {
	Id       sql.NullInt64
	Username sql.NullString
	Score    sql.NullInt16
}
