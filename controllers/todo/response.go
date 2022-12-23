package todo

import "time"

type createTodo struct {
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ID         uint      `json:"id"`
	ActivityID uint      `json:"activity_group_id"`
	Title      string    `json:"title"`
	IsActive   bool      `json:"is_active"`
	Priority   string    `json:"priority"`
}

type resultResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
