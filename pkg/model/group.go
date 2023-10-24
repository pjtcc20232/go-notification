package model

import (
	"encoding/json"

	"github.com/notification/back-end/internal/config/logger"
)

type Group struct {
	ID         uint64 `json:"_id"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
}

type GroupList struct {
	List []*Group `json:"list"`
}

func (cl GroupList) String() string {
	data, err := json.Marshal(cl)

	if err != nil {
		logger.Error("error to convert GroupList to JSON:"+err.Error(), err)

		return ""
	}

	return string(data)
}
