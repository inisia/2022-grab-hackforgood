package models

import (
	"net/http"
	"time"
)

type MetaOrder struct {
	Id          string    `json:"id"`
	PartnerId   string    `json:"partnerId"`
	Status      string    `json:"status"`
	StartAt     time.Time `json:"startAt"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Orders      []*Order  `json:"orders"`
}

func (rd *MetaOrder) Render(w http.ResponseWriter, r http.Request) error {
	return nil
}
