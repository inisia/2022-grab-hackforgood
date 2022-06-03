package partners

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/inisia/grab-as-a-service/backend/db/conn"
	"github.com/inisia/grab-as-a-service/backend/db/daos"
)

func GetOwnDetail(w http.ResponseWriter, r *http.Request) {
	partnerId := r.Header.Get("X-Partner-Id")
	conn := conn.CreateDBConnection()
	partnerDao := daos.PartnerDao{Conn: conn}
	partner := partnerDao.Detail(partnerId)

	render.JSON(w, r, partner)
}
