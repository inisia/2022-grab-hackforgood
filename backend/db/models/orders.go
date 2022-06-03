package models

import (
	"net/http"
	"time"
)

type Order struct {
	Id          string     `json:"id"`
	MetaOrderId string     `json:"metaOrderId"`
	MetaOrder   *MetaOrder `json:"metaorder"`
	StartAt     time.Time  `json:"startAt"`
	FromLang    float64    `json:"fromLang"`
	FromLat     float64    `json:"fromLat"`
	ToLang      float64    `json:"toLang"`
	ToLat       float64    `json:"toLat"`
	Status      string     `json:"status"`
	ClientName  string     `json:"clientName"`
	Contact     string     `json:"contact"`
	Price       float64    `json:"price"`
}

func (rd *Order) Render(w http.ResponseWriter, r http.Request) error {
	return nil
}
