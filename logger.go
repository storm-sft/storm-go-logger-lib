package storm_logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"github.com/storm-sft/storm-go-logger-lib/env"
)

// LOGGER is the global logger used in the application.
var LOGGER *zap.SugaredLogger

// InitLogger initializes the logger used in the application.
// It creates a new logger using the zap library and the provided mode.
// If the logger is successfully created, it assigns the logger to the global LOGGER variable.
// If the logger cannot be created, it logs the error and exits the program.
func InitLogger(environment env.Environment) {
	var logger *zap.Logger
	var err error
	if environment == env.Production {
		logger, err = zap.NewProduction()
	} else {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.FunctionKey = "func"
		logger, err = config.Build()
	}

	if err != nil {
		// Exits the program
		log.Fatal("logger initialization failed", err)
	}

	LOGGER = logger.Sugar()
	LOGGER.Info("logger successfully initialized")
}
