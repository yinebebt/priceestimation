package persist

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yinebebt/priceestimation/internal/constants/model/db"
	"github.com/yinebebt/priceestimation/platform"
)

type DB struct {
	*db.Queries
	pool *pgxpool.Pool
	log  platform.Logger
}

func New(pool *pgxpool.Pool, log platform.Logger) DB {
	return DB{
		Queries: db.New(pool),
		pool:    pool,
		log:     log,
	}
}
