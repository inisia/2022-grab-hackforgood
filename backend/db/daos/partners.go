package daos

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"

	"github.com/inisia/grab-as-a-service/backend/db/models"
)

type PartnerDao struct {
	Conn *pgx.Conn
}

func (state *PartnerDao) Detail(id string) *models.Partner {
	partner := &models.Partner{}

	row := state.Conn.QueryRow(context.Background(), "SELECT * FROM partners WHERE id=$1;", id)

	err := row.Scan(&partner.Id, &partner.Name, &partner.Quota)

	if err != nil {
		panic(err)
	}

	return partner
}

func (state *PartnerDao) AddQuota(id string, quota int64) *models.Partner {
	log.Println(id, quota)
	_, err := state.Conn.Exec(context.Background(), "UPDATE partners SET quota = quota + $1 WHERE ID = $2", quota, id)

	if err != nil {
		panic(err)
	}

	return state.Detail(id)
}
