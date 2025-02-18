package web_question

type Result struct {
	Result []ItemResult `json:"result"`
}

type ItemResult struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
}
