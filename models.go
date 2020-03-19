package main

type UserQuery struct {
	UserId string `json:"user_id"`
	Text   string `json:"text"`
	//DueDate time.Time `json:"due_date"`
}
