package logger

import (
    "fmt"
    "net/http"
    "time"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"

    "github.com/2pizzzza/IskenderBackend/internal/config"
)

const (
    envDebug = "debug"
    envLocal = "dev"
    envProd  = "prod"
)

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}

func New(config *config.Config) (*zap.Logger, error) {
    var cfg zap.Config

    consoleEncoderConfig := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "msg",
        StacktraceKey:  "",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.CapitalColorLevelEncoder, 
        EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
            enc.AppendString(t.Format("2006-01-02 15:04:05"))
        },
        EncodeDuration: zapcore.StringDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    }

    jsonEncoderConfig := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.CapitalLevelEncoder, 
        EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
            enc.AppendString(t.Format("2006-01-02 15:04:05"))
        },
        EncodeDuration: zapcore.StringDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    }

    switch config.Log.Level {
    case envDebug:
        cfg = zap.NewDevelopmentConfig()
        cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
        cfg.Encoding = "console"
        cfg.EncoderConfig = consoleEncoderConfig
    case envLocal:
        cfg = zap.NewDevelopmentConfig()
        cfg.Encoding = "console"
        cfg.EncoderConfig = consoleEncoderConfig
        cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
    case envProd:
        cfg = zap.NewProductionConfig()
        cfg.Encoding = "json"
        cfg.EncoderConfig = jsonEncoderConfig
        cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
    default:
        return nil, fmt.Errorf("invalid log level: %s", config.Log.Level)
    }

    logger, err := cfg.Build()
    if err != nil {
        return nil, fmt.Errorf("failed to initialize logger: %w", err)
    }

    return logger, nil
}

func LoggingMiddleware(log *zap.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()

            rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

            next.ServeHTTP(rw, r)

            isProd := log.Core().Enabled(zapcore.InfoLevel) && !log.Core().Enabled(zapcore.DebugLevel)

            statusText := formatStatus(rw.statusCode, isProd)

            msg := fmt.Sprintf("%s %s %s %s %s",
                r.Method,
                r.URL.Path,
                statusText,
                r.RemoteAddr,
                time.Since(start),
            )

            if rw.statusCode >= 500 {
                log.Error(msg)
            } else if rw.statusCode >= 400 {
                log.Warn(msg)
            } else {
                log.Info(msg)
            }
        })
    }
}

func formatStatus(status int, isProd bool) string {
    statusText := fmt.Sprintf("%d", status)
    if !isProd {
        switch {
        case status >= 200 && status < 300:
            statusText = fmt.Sprintf("\x1b[32m%d\x1b[0m", status) 
        case status >= 400 && status < 500:
            statusText = fmt.Sprintf("\x1b[33m%d\x1b[0m", status) 
        case status >= 500:
            statusText = fmt.Sprintf("\x1b[31m%d\x1b[0m", status) 
        }
    }
    return statusText
}