package conn

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
)

var _conn *pgx.Conn

func CreateDBConnection() *pgx.Conn {

	if _conn != nil {
		return _conn
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	_conn = conn

	return _conn
}
