package utils

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strings"
)

const envLogLevel = "LOG_LEVEL"

func InitLogger() {
	configJson := fmt.Sprintf(`{
		"level": "%s",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "timeKey": "time",
		  "levelEncoder": "lowercase",
		  "timeEncoder": "iso8601"
		}
	  }`, getLevel())
	rawJSON := []byte(configJson)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	logger := zap.Must(cfg.Build())
	zap.ReplaceGlobals(logger)

	defer logger.Sync()
}

func getLevel() string {
	return strings.ToLower(strings.TrimSpace(os.Getenv(envLogLevel)))
}

func LoggerInfo(message string, tags ...zap.Field) {
	zap.L().Info(message, tags...)
}

func LoggerError(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	zap.L().Error(message, tags...)
}

func LoggerPanic(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	zap.L().Fatal(message, tags...)
}
