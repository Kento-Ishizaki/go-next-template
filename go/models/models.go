package models

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Email string `json:"email"`
}
