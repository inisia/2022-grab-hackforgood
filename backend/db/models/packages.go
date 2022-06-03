package models

import (
	"net/http"
)

type Package struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Quota       int64  `json:"quota"`
}

func (rd *Package) Render(w http.ResponseWriter, r http.Request) error {
	return nil
}
