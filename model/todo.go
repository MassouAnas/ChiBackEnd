package model

type Todo struct {
	TodoID     uint64  `json:"todo_id"`
	TodoTitle  string  `json:"todo_title"`
	TodoDescr  string  `json:"todo_description"`
	CreatedAt  *string `json:"created_at"`
	FinishedAt string  `json:"finished_at"`
}