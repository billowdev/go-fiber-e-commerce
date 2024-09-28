package logs

import "go.uber.org/zap"

func NewLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.Encoding = "json" // Use JSON encoding for structured logs
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	return logger
}
