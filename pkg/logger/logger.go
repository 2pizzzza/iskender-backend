package logger

import (
    "fmt"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"

    "github.com/2pizzzza/IskenderBackend/internal/config"
)

const (
    envDebug = "debug"
    envLocal = "dev"
    envProd  = "prod"
)

func New(config *config.Config) (*zap.Logger, error) {
    var cfg zap.Config
	
    switch config.Log.Level {
    case envDebug:
        cfg = zap.NewDevelopmentConfig()
        cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
    case envLocal:
        cfg = zap.NewDevelopmentConfig()
        cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
    case envProd:
        cfg = zap.NewProductionConfig()
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