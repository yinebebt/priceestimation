package initiator

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yinebebt/priceestimation/platform/logger"
)

func InitDB(url string, log logger.Logger) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("Failed to connect to database: %v", err))
	}
	config.ConnConfig.Logger = log.Named("pgx")
	config.MaxConns = 1000
	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("Failed to connect to database: %v", err))
	}
	// Check the connection by executing a query
	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal(context.Background(), err.Error())
	}

	return conn
}
