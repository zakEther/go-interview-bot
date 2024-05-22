package entities

import "encoding/json"

type Question struct {
	QuestionID      int      `json:"question_id"`
	QuestionText    string   `json:"question_text"`
	QuestionOptions []string `json:"question_options"`
	Answer          int      `json:"answer"`
	Explanation     string   `json:"explanation"`
}

func (q *Question) MarshalJSON() ([]byte, error) {
	type Alias Question
	return json.Marshal(&struct {
		*Alias
		QuestionOptions []string `json:"question_options"`
	}{
		Alias:           (*Alias)(q),
		QuestionOptions: q.QuestionOptions,
	})
}

func (q *Question) UnmarshalJSON(data []byte) error {
	type Alias Question
	aux := &struct {
		*Alias
		QuestionOptions []string `json:"question_options"`
	}{
		Alias: (*Alias)(q),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	q.QuestionOptions = aux.QuestionOptions
	return nil
}

func (q *Question) GetID() int {
	return q.QuestionID
}

func (q *Question) GetText() string {
	return q.QuestionText
}

func (q *Question) GetQuestionOptions() []string {
	return q.QuestionOptions
}

func (q *Question) GetAnswer() int {
	return q.Answer
}
