# Logging Package Detailed Explanation

## Key Changes and Explanations

### Dual Output (Console and File):

```go
core := zapcore.NewTee(consoleCore, fileCore)
```


- `zapcore.NewTee` is used to combine two logging cores: one for the console and one for the file.
- This allows logs to be written to both outputs simultaneously, ensuring logs are captured in both places.

### Encoders:

```go

consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
```

- `consoleEncoder` uses a console-friendly format, which is more readable for humans.
- `fileEncoder` uses a JSON format, which is structured and suitable for log analysis tools.
- Both encoders are customizable based on your logging requirements.

### Logging Level:

```go
logLevel := zapcore.DebugLevel
```

- The logging level is set to `zapcore.DebugLevel`, which means all log levels (`Debug`, `Info`, `Warn`, `Error`) will be captured.
- This can be adjusted if you want to capture fewer logs (e.g., starting at `Info` level or higher).

### Options:

```go
logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
```

- `zap.AddCaller()` adds the file name and line number of the caller to the logs, which is useful for debugging.
- `zap.AddStacktrace(zapcore.ErrorLevel)` captures stack traces for errors, making it easier to trace issues in the code.

## General logs

```go
// Simple informational log
logs.Info("Service started successfully")

// Informational log with additional context
logs.Info("Service started successfully", zap.String("service", "auth"), zap.String("version", "1.0.0"))

// Error log with additional context
err := someFunction()
if err != nil {
    logs.Error(err, "Failed to process request", zap.String("userID", "12345"))
}

// Warning log with context
logs.Warning("Cache miss", zap.String("cacheKey", "user_12345"))

// Feature-specific info log
logs.InfoHandler("UserManagement", "User login succeeded", zap.String("userID", "12345"))
```



## HTTP Logs

Usage Example
To use the separate loggers for request and response logging, you can use httpLogger for HTTP-related logs:

```go
// Example of logging an HTTP request
func LogHTTPRequest(method, url string, statusCode int) {
    httpLogger.Info(
        "HTTP Request",
        zap.String("method", method),
        zap.String("url", url),
        zap.Int("status_code", statusCode),
    )
}

// Example of logging an HTTP response
func LogHTTPResponse(statusCode int, responseBody string) {
    httpLogger.Info(
        "HTTP Response",
        zap.Int("status_code", statusCode),
        zap.String("response_body", responseBody),
    )
}
```

### Summary
- Separate Loggers: Create separate loggers for different log types (e.g., general logs and HTTP logs).
- Separate Files: Direct logs to different files using separate cores and write syncers.
- Encoders: Use the same or different encoders based on your needs for each type of log.