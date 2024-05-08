package model

import "time"

type Todo struct {
	TodoID    uint64 `json:"todo_id"`
	TodoTitle string `json:"todo_title"`
	TodoDescr string `json:"todo_description"`
	CreatedAt *time.Time `json:"created_at"`
	FinishedAt *time.Time `json:"finished_at"`
}