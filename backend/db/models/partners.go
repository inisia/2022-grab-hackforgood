package models

import (
	"net/http"
)

type Partner struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Quota int64  `json:"quota"`
}

func (rd *Partner) Render(w http.ResponseWriter, r http.Request) error {
	return nil
}
