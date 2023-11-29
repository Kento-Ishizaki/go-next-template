package entity

type Todo struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID int    `json:"userId" gorm:"column:user_id"`
}
