package initiator

import (
	"context"
	"errors"
	"fmt"
	"github.com/yinebebt/priceestimation/platform/logger"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"go.uber.org/zap"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitiateMigration(path, conn string, log logger.Logger) *migrate.Migrate {
	conn = fmt.Sprintf("postgres://%s", strings.Split(conn, "://")[1])
	m, err := migrate.New(fmt.Sprintf("file://%s", path), conn)
	if err != nil {
		log.Fatal(context.Background(), "could not create migrator", zap.Error(err))
	}
	return m
}

func UpMigration(m *migrate.Migrate, log logger.Logger) {
	err := m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(context.Background(), "could not migrate", zap.Error(err))
	}
}
