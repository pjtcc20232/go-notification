package model

type Student struct {
	ID           uint64 `json:"_id"`
	Name         string `json:"name"`
	Registration string `json:"registration"`
	Tbl_class_id uint64 `json:"id_class"`
	Tbl_usr_id   uint64 `json:"id_usr"`
}
