package main

import "time"

type Wish struct {
	EmployeeID string    `json:"employee_id"`
	Keywords   string    `json:"keywords"`
	CreatedAt  time.Time `json:"created_at"`
}
