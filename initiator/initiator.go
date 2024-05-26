package initiator

import (
	"context"
	"fmt"
	"github.com/yinebebt/priceestimation/internal/constants/query/persist"
	"github.com/yinebebt/priceestimation/platform/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Initiator
// @title price-estimation api
// @version         0.1.0
// @contact.name   Support Email
// @contact.url    Contact_url
// @contact.email  contact@price_estimation
// @host localhost
// @BasePath  /api/v1
func Initiator(ctx context.Context) {
	log := logger.New(InitLogger())
	log.Info(ctx, "logger initialized")

	log.Info(ctx, "initializing config")
	configName := "config"
	if name := os.Getenv("CONFIG_NAME"); name != "" {
		configName = name
		log.Info(ctx, fmt.Sprintf("config name is set to %s", configName))
	} else {
		log.Info(ctx, "using default config name 'config'")
	}
	InitConfig(configName, "config", log)
	log.Info(ctx, "config initialized")

	log.Info(ctx, "initializing database")
	Conn := InitDB(viper.GetString("database.url"), log)
	log.Info(ctx, "database initialized")

	if viper.GetBool("database.migration.active") {
		log.Info(context.Background(), "initializing migration")
		m := InitiateMigration(viper.GetString("database.migration.path"), viper.GetString("database.url"), log)
		UpMigration(m, log)
		log.Info(context.Background(), "migration initialized")
	}
	// if there, do platform initialization, and database policies setup here
	//
	log.Info(ctx, "initializing persistence layer")
	persistence := InitPersistence(persist.New(Conn, log), log)
	log.Info(ctx, "persistence layer initialized")

	// if there, initialize state here

	log.Info(ctx, "initializing module")
	_ = InitModule(persistence, log)
	log.Info(ctx, "module initialized")

	log.Info(ctx, "initializing server")
	server := gin.New()

	log.Info(ctx, "server initialized")

	log.Info(ctx, "initializing router")
	_ = server.Group("/api/v1")

	log.Info(ctx, "router initialized")

	srv := &http.Server{
		Addr:              viper.GetString("server.host") + ":" + viper.GetString("server.port"),
		ReadHeaderTimeout: viper.GetDuration("read_header_timeout"),
		Handler:           server,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	go func() {
		log.Info(ctx, "server started",
			zap.String("host", viper.GetString("server.host")),
			zap.Int("port", viper.GetInt("server.port")))
		log.Info(ctx, fmt.Sprintf("server stopped with error %v", srv.ListenAndServe()))
	}()
	sig := <-quit
	log.Info(ctx, fmt.Sprintf("server shutting down with signal %v", sig))
	ctx, cancel := context.WithTimeout(ctx, viper.GetDuration("server.timeout"))
	defer cancel()

	log.Info(ctx, "shutting down server")
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(context.Background(), fmt.Sprintf("error while shutting down server: %v", err))
	}
	log.Info(context.Background(), "server shutdown complete")
}
