package model

import (
	"encoding/json"

	"github.com/notification/back-end/internal/config/logger"
)

type Class struct {
	ID            uint64 `json:"_id"`
	Schedules     string `json:"schedules"`
	Tbl_course_id int    `json:"_tbl_curso_id"`
}

type ClasstList struct {
	List []*Class `json:"list"`
}

func (cl ClasstList) String() string {
	data, err := json.Marshal(cl)

	if err != nil {
		logger.Error("error to convert ClassList to JSON:"+err.Error(), err)

		return ""
	}

	return string(data)
}
