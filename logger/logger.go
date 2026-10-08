package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.Logger {
	file, err := os.OpenFile(
		"./logs/app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		panic(err)
	}
	writeSyncer := zapcore.AddSync(file)
	encoder := zapcore.NewJSONEncoder(
		zap.NewProductionEncoderConfig(),
	)
	core := zapcore.NewCore(
		encoder,
		writeSyncer,
		zap.InfoLevel,
	)
	return zap.New(core)
}
