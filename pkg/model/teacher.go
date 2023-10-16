package model

type Teacher struct {
	ID         uint64 `json:"_id"`
	Name       string `json:"name"`
	Chair      uint64 `json:"Chair"`
	Tbl_usr_id string `json:"id_usr"`
}
