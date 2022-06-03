package daos

import (
	"context"

	"github.com/jackc/pgx/v4"

	"github.com/inisia/grab-as-a-service/backend/db/models"
)

type MetaOrderDao struct {
	Conn *pgx.Conn
}

func (state *MetaOrderDao) List() []*models.MetaOrder {
	list := []*models.MetaOrder{}

	rows, err := state.Conn.Query(context.Background(), "SELECT id, partner_id, start_at, description, status, price FROM meta_orders;")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		metaOrder := &models.MetaOrder{}
		err := rows.Scan(&metaOrder.Id, &metaOrder.PartnerId, &metaOrder.StartAt, &metaOrder.Description, &metaOrder.Status, &metaOrder.Price)

		if err != nil {
			panic(err)
		}

		list = append(list, metaOrder)
	}

	return list
}

func (state *MetaOrderDao) Detail(id string) *models.MetaOrder {
	metaOrder := &models.MetaOrder{}

	row := state.Conn.QueryRow(context.Background(), "SELECT id, start_at, partner_id, description, status, price FROM meta_orders WHERE id=$1;", id)
	err := row.Scan(&metaOrder.Id, &metaOrder.StartAt, &metaOrder.PartnerId, &metaOrder.Description, &metaOrder.Status, &metaOrder.Price)

	if err != nil {
		panic(err)
	}

	query := `
    SELECT id, meta_order_id, start_at, from_lang, from_lat, to_lang, to_lat, status, client_name, contact, price
    FROM orders WHERE meta_order_id=$1;
  `
	rows, err := state.Conn.Query(context.Background(), query, id)

	if err != nil {
		panic(err)
	}

	orders := []*models.Order{}

	for rows.Next() {
		order := &models.Order{}
		err := rows.Scan(&order.Id, &order.MetaOrderId, &order.StartAt, &order.FromLang,
			&order.FromLat, &order.ToLang, &order.ToLat, &order.Status, &order.ClientName, &order.Contact, &order.Price,
		)

		if err != nil {
			panic(err)
		}

		orders = append(orders, order)
	}

	metaOrder.Orders = orders

	return metaOrder
}

func (state *MetaOrderDao) Create(metaOrder models.MetaOrder) *models.MetaOrder {
	tx, err := state.Conn.Begin(context.Background())

	if err != nil {
		panic(err)
	}

	query := `
    INSERT INTO
    meta_orders (partner_id, status, start_at, description, price) 
    VALUES ($1, $2, $3, $4, $5)
    RETURNING
    id
  `
	err = tx.QueryRow(context.Background(), query,
		metaOrder.PartnerId, "wait", metaOrder.StartAt, metaOrder.Description, metaOrder.Price,
	).Scan(&metaOrder.Id)

	if err != nil {
		panic(err)
	}

	for _, order := range metaOrder.Orders {
		query = `
      INSERT INTO
      orders (meta_order_id, start_at, from_lang, from_lat, to_lang, to_lat, status, client_name, contact, price)
      VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    `

		_, err = tx.Exec(context.Background(), query,
			metaOrder.Id,
			order.StartAt, order.FromLang, order.FromLat,
			order.ToLang, order.ToLat, "wait",
			order.ClientName, order.Contact, order.Price,
		)

		if err != nil {
			panic(err)
		}
	}

	err = tx.Commit(context.Background())

	if err != nil {
		panic(err)
	}
	return state.Detail(metaOrder.Id)
}
