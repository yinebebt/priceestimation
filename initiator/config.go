package initiator

import (
	"context"
	"fmt"
	"github.com/yinebebt/priceestimation/platform/logger"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig(name, path string, log logger.Logger) {
	viper.SetConfigFile(fmt.Sprintf("%s/%s.yaml", path, name))
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(context.Background(), fmt.Sprintf("Failed to read config: %v", err))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Panic(context.Background(), "Config file changed:", zap.String("file", e.Name))
	})
}

func GetMapSlice(path string) []map[string]any {
	value := viper.Get(path)
	mapInterfaceSlice, ok := value.([]any)
	if !ok {
		return nil
	}

	var mapStringAny []map[string]any
	for _, v := range mapInterfaceSlice {
		v, ok := v.(map[string]any)
		if ok {
			mapStringAny = append(mapStringAny, v)
		}
	}
	return mapStringAny
}
