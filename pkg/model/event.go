package model

import "time"

type Event struct {
	ID                uint64    `json:"_id"`
	Name              string    `json:"name"`
	EventDate         time.Time `json:"event_date"`
	Description       string    `json:"description"`
	Tbl_class_id      uint64    `json:"id_class"`
	Tbl_class_teacher uint64    `json:"id_teacher"`
	StatusEvent       string    `json:"status_evento"`
	CreatedAt         string    `json:"created_at,omitempty"`
	UpdatedAt         string    `json:"updated_at,omitempty"`
}
