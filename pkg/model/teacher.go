package model

import (
	"encoding/json"

	"github.com/notification/back-end/internal/config/logger"
)

type Teacher struct {
	ID         uint64 `json:"_id"`
	Name       string `json:"name"`
	Chair      uint64 `json:"Chair"`
	Tbl_usr_id string `json:"id_usr"`
}

type TeacherList struct {
	List []*Teacher `json:"list"`
}

func (tl TeacherList) String() string {
	data, err := json.Marshal(tl)

	if err != nil {
		logger.Error("error to convert ClassList to JSON:"+err.Error(), err)

		return ""
	}

	return string(data)
}
