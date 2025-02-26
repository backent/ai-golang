package web_exam

type ExamSubmitRequest struct {
	Id          int64        `json:"id"`
	Submissions []Submission `json:"submissions"`
}

type Submission struct {
	Question    string   `json:"question"`
	Answer      string   `json:"answer"`
	UserAnswer  string   `json:"user_answer"`
	Options     []string `json:"options"`
	Explanation string
}
