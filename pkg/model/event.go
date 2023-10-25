package model

import (
	"encoding/json"
	"time"

	"github.com/notification/back-end/internal/config/logger"
)

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

type EventList struct {
	List []*Event `json:"list"`
}

func (ev GroupList) String() string {
	data, err := json.Marshal(ev)

	if err != nil {
		logger.Error("error to convert GroupList to JSON:"+err.Error(), err)

		return ""
	}

	return string(data)
}
