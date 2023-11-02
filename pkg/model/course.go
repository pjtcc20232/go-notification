package model

import (
	"encoding/json"

	"github.com/notification/back-end/internal/config/logger"
)

type Courses struct {
	ID   int    `json:"_id"`
	Name string `json:"name"`
}

type CourseList struct {
	List []*Courses `json:"list"`
}

func (cl CourseList) String() string {
	data, err := json.Marshal(cl)

	if err != nil {
		logger.Error("error to convert CourseList to JSON:"+err.Error(), err)

		return ""
	}

	return string(data)
}
