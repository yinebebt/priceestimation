package initiator

import (
	"go.uber.org/zap"
	"log"
	"os"
)

func InitLogger() *zap.Logger {
	logFilePath := "platform/logger/logs.log"
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		file, err := os.Create(logFilePath)
		if err != nil {
			log.Printf("failed to create log file: %v", err)
		}
		err = file.Close()
		if err != nil {
			log.Printf("failed to close log file: %v", err)
		}
	}

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stderr", logFilePath}
	logger, err := cfg.Build() // zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	return logger
}
