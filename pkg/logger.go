package pkg

import (
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLog() {
	now := time.Now()
	logPath := filepath.Join("storage/logs", now.Format("2006-01")+".log")
	_ = os.MkdirAll(filepath.Dir(logPath), os.ModePerm)

	lumberjackLogger := &lumberjack.Logger{
		Filename: logPath,
		MaxSize:  10,
	}

	writeSyncer := zapcore.AddSync(lumberjackLogger)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		writeSyncer,
		zapcore.InfoLevel,
	)

	Logger = zap.New(core)
}
