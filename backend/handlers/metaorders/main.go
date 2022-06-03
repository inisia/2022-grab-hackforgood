package metaorders

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/go-chi/render"

	"github.com/inisia/grab-as-a-service/backend/db/conn"
	"github.com/inisia/grab-as-a-service/backend/db/daos"
	"github.com/inisia/grab-as-a-service/backend/db/models"
)

func injectPrice(metaOrder *models.MetaOrder) {
	for _, order := range metaOrder.Orders {
		degree := math.Pow(math.Pow(order.ToLang-order.FromLang, 2.0)+math.Pow(order.ToLat-order.FromLat, 2.0), 0.5)
		price := degree * 111.0 * 6000.0
		order.Price = price
		metaOrder.Price += order.Price
	}
}

func GetMetaOrders(w http.ResponseWriter, r *http.Request) {
	db := conn.CreateDBConnection()
	packageDao := daos.MetaOrderDao{Conn: db}
	list := packageDao.List()

	render.JSON(w, r, list)
}

func CreateMetaOrder(w http.ResponseWriter, r *http.Request) {
	db := conn.CreateDBConnection()
	metaorderDao := daos.MetaOrderDao{Conn: db}

	decoder := json.NewDecoder(r.Body)

	body := &models.MetaOrder{}
	err := decoder.Decode(&body)

	if err != nil {
		panic(err)
	}

	body.PartnerId = r.Header.Get("X-Partner-Id")
	injectPrice(body)

	metaOrder := metaorderDao.Create(*body)

	render.JSON(w, r, metaOrder)
}
