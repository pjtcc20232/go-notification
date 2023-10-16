package model

type Group struct {
	ID         uint64 `json:"_id"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
}
