package models

import (
	"database/sql"
)

var ExamTable string = "exams"

type Exam struct {
	Id          int64
	Username    string
	QuestionId  int64
	Score       int16
	Submissions string
	ExamAt      sql.NullTime
}
