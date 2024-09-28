package logs

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger     *zap.Logger
	httpLogger *zap.Logger
)

func init() {
	// Create an encoder config that defines how logs will be formatted
	encoderConfig := zapcore.EncoderConfig{
		LevelKey:    "level",
		TimeKey:     "timestamp",
		MessageKey:  "message",
		EncodeLevel: zapcore.LowercaseLevelEncoder, // e.g., "debug", "info", "warn", "error"
		EncodeTime:  zapcore.ISO8601TimeEncoder,    // e.g., "2024-09-09T14:00:00Z"
	}

	// Open the log file for writing
	logFile, err := os.OpenFile("logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	// Ensure the log file is closed when the application exits
	// defer logFile.Close() // Remove if you want the file to stay open for the lifetime of the app

	// Set up the HTTP log file
	httpLogFile, err := os.OpenFile("http_logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open HTTP log file: %v", err)
	}

	// Create JSON encoders for both the console and file logs
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// HTTP logs encoder (can be different if needed)
	httpConsoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	httpFileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	httpFileWriteSyncer := zapcore.AddSync(httpLogFile)

	// Create write syncers to define where the logs will be written
	consoleWriteSyncer := zapcore.Lock(os.Stdout)
	fileWriteSyncer := zapcore.AddSync(logFile)

	// Set the log level to debug (logs all levels: debug, info, warn, error)
	logLevel := zapcore.DebugLevel

	// General log core
	// Combine the console and file cores so that logs are written to both
	consoleCore := zapcore.NewCore(consoleEncoder, consoleWriteSyncer, logLevel)
	fileCore := zapcore.NewCore(fileEncoder, fileWriteSyncer, logLevel)
	core := zapcore.NewTee(consoleCore, fileCore)

	// Create the logger with the combined core
	logger = zap.New(core)

	// HTTP log core
	httpConsoleCore := zapcore.NewCore(httpConsoleEncoder, consoleWriteSyncer, logLevel)
	httpFileCore := zapcore.NewCore(httpFileEncoder, httpFileWriteSyncer, logLevel)
	httpCore := zapcore.NewTee(httpConsoleCore, httpFileCore)
	httpLogger = zap.New(httpCore)

	// Optionally, add stacktrace and caller info to the logs
	logger = logger.WithOptions(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	httpLogger = httpLogger.WithOptions(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

}

// Info logs informational messages to both the log file and the console.
func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}


// Error logs error messages with the error and additional context.
func Error(err error, message string, fields ...zap.Field) {
	log.Printf("Error: %s - %v\n", message, err)
	allFields := append(fields, zap.Error(err))
	logger.Error(message, allFields...)
}

// Warning logs warning messages to both the log file and the console.
func Warning(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}
