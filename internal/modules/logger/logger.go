package logger

import "go.uber.org/zap"

func New() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()

	if err != nil {
		return nil, err
	}

	return logger, nil
}

func Sugarize(logger *zap.Logger) *zap.SugaredLogger {
	return logger.Sugar()
}
