package daos

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/inisia/grab-as-a-service/backend/db/models"
)

type PackageDao struct {
	Conn *pgx.Conn
}

func (state *PackageDao) List() []*models.Package {
	list := []*models.Package{}

	rows, err := state.Conn.Query(context.Background(), "SELECT * FROM packages;")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		pkg := &models.Package{}
		err := rows.Scan(&pkg.Id, &pkg.Name, &pkg.Description, &pkg.Price, &pkg.Quota)

		if err != nil {
			fmt.Println(err)
		}

		list = append(list, pkg)
	}

	return list
}

func (state *PackageDao) Detail(id string) *models.Package {
	pkg := &models.Package{}

	row := state.Conn.QueryRow(context.Background(), "SELECT * FROM packages WHERE id=$1;", id)

	err := row.Scan(&pkg.Id, &pkg.Name, &pkg.Description, &pkg.Price, &pkg.Quota)

	if err != nil {
		panic(err)
	}

	return pkg
}
