package AttProgram

import "time"
type TrainingRecord struct {
	ID           string    `json:"id"`
	TrainingDate time.Time `json:"training_date"`
	Comment string `json:"comment"`
	TrainingResult Results `json:"training_results"`
}

type ShareRequest struct {
	Author User `json:"author"`
	TrainingResult TrainingRecord `json:"training_result"`
	ShareTo []User `json:"share_to"`
}

type  User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	TrainingResults []TrainingRecord `json:"training_results"`
}

type  Results struct {
	BurnedCalories int `json:"burned_calories"`
	TrainingDuration time.Duration `json:"training_duration"`
	Exercises []string `json:"exercises"`
}
