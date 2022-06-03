package packages

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/inisia/grab-as-a-service/backend/db/conn"
	"github.com/inisia/grab-as-a-service/backend/db/daos"
)

func GetPackages(w http.ResponseWriter, r *http.Request) {
	db := conn.CreateDBConnection()
	packageDao := daos.PackageDao{Conn: db}
	list := packageDao.List()

	render.JSON(w, r, list)
}

func BuyPackage(w http.ResponseWriter, r *http.Request) {
	db := conn.CreateDBConnection()
	packageDao := daos.PackageDao{Conn: db}
	partnerDao := daos.PartnerDao{Conn: db}

	packageId := chi.URLParam(r, "packageId")
	partnerId := r.Header.Get("X-Partner-Id")

	pkg := packageDao.Detail(packageId)
	partner := partnerDao.AddQuota(partnerId, pkg.Quota)

	render.JSON(w, r, partner)
}
