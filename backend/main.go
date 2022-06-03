package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/inisia/grab-as-a-service/backend/db/conn"
	"github.com/inisia/grab-as-a-service/backend/handlers/metaorders"
	"github.com/inisia/grab-as-a-service/backend/handlers/packages"
	"github.com/inisia/grab-as-a-service/backend/handlers/partners"
)

func main() {

	db := conn.CreateDBConnection()
	defer db.Close(context.Background())

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/packages", packages.GetPackages)
	r.Post("/packages/{packageId}/buy", packages.BuyPackage)

	r.Get("/partners/me", partners.GetOwnDetail)

	r.Post("/metaorders/", metaorders.CreateMetaOrder)

	http.ListenAndServe(":3333", r)
}
