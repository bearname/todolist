package model

type Task struct {
	Id            string  `json:"id"`
	Description   string  `json:"description"`
	Status        bool    `json:"status"`
	CreatedDate   []uint8 `json:"created_date"`
	CompletedDate []uint8 `json:"completed_date"`
}
