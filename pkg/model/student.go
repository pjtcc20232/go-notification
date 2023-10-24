package model

import (
	"encoding/json"

	"github.com/notification/back-end/internal/config/logger"
)

type Student struct {
	ID           uint64 `json:"_id"`
	Name         string `json:"name"`
	Registration string `json:"registration"`
	Period       string `json:"period"`
	Tbl_class_id uint64 `json:"id_class"`
	Tbl_usr_id   uint64 `json:"id_usr"`
}

type StudenttList struct {
	List []*Student `json:"list"`
}

func (sl StudenttList) String() string {
	data, err := json.Marshal(sl)

	if err != nil {
		logger.Error("error to convert ClassList to JSON:"+err.Error(), err)

		return ""
	}

	return string(data)
}
