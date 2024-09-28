package logs

import "go.uber.org/zap"

func LogSysInfo(message, serviceName, methodName string, additionalFields ...zap.Field) {
	fields := []zap.Field{
		zap.String("service", serviceName),
		zap.String("method", methodName),
	}

	fields = append(fields, additionalFields...)

	Info(message, fields...)
}

// LogInfo is a higher-level function that logs informational messages with standard fields such as trace ID, user ID, status code, service name, and method name.
// It also accepts additional zap fields for more detailed logging.
func LogInfo(message, traceID, userID, statusCode, serviceName, methodName string, additionalFields ...zap.Field) {
	// Create a slice to hold all fields
	fields := []zap.Field{
		zap.String("trace_id", traceID),
		zap.String("uid", userID),
		zap.String("code", statusCode),
		zap.String("service", serviceName),
		zap.String("method", methodName),
	}

	// Append additional fields
	fields = append(fields, additionalFields...)

	// Pass all fields to the Info function
	Info(message, fields...)
}

func LogSysWarning(message, code, serviceName, methodName string, additionalFields ...zap.Field) {
	fields := []zap.Field{
		zap.String("code", code),
		zap.String("service", serviceName),
		zap.String("method", methodName),
	}

	fields = append(fields, additionalFields...)

	Warning(message, fields...)
}

// LogWarning logs a warning message with standard fields such as trace ID, user ID, service name, and method name.
func LogWarning(message, traceID, userID, code, serviceName, methodName string, additionalFields ...zap.Field) {
	// Create a slice to hold all fields
	fields := []zap.Field{
		zap.String("trace_id", traceID),
		zap.String("uid", userID),
		zap.String("code", code),
		zap.String("service", serviceName),
		zap.String("method", methodName),
	}

	// Append additional fields
	fields = append(fields, additionalFields...)

	// Pass all fields to the Warning function
	Warning(message, fields...)
}

// LogSysError logs an error message with the error object, service name, and method name.
func LogSysError(err error, message, code, serviceName, methodName string, additionalFields ...zap.Field) {
	// Create a slice to hold all fields
	fields := []zap.Field{
		zap.String("code", code),
		zap.String("service", serviceName),
		zap.String("method", methodName),
	}

	// Append additional fields
	fields = append(fields, additionalFields...)

	// Pass all fields to the Error function
	Error(err, message, fields...)
}

// LogError logs an error message with the error object, trace ID, user ID, status code, service name, and method name.
func LogError(err error, message, traceID, userID, code, serviceName, methodName string, additionalFields ...zap.Field) {
	// Create a slice to hold all fields
	fields := []zap.Field{
		zap.String("trace_id", traceID),
		zap.String("uid", userID),
		zap.String("code", code),
		zap.String("service", serviceName),
		zap.String("method", methodName),
	}

	// Append additional fields
	fields = append(fields, additionalFields...)

	// Pass all fields to the Error function
	Error(err, message, fields...)
}
