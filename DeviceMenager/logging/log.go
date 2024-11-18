package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log Logger

func init() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	sugar := logger.Sugar()

	Log = &loggerWrapper{lw: sugar}
}

type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info( args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

type loggerWrapper struct {
	lw *zap.SugaredLogger
}
func (logger *loggerWrapper) Errorf(format string, args ...interface{}) {
	logger.lw.Errorf(format, args)
}
func (logger *loggerWrapper) Fatalf(format string, args ...interface{}) {
	logger.lw.Fatalf(format, args)
}
func (logger *loggerWrapper) Fatal(args ...interface{}) {
	logger.lw.Fatal(args)
}
func (logger *loggerWrapper) Info(args ...interface{}) {
	logger.lw.Info(args...)
}
func (logger *loggerWrapper) Infof(format string, args ...interface{}) {
	logger.lw.Infof(format, args)
}
func (logger *loggerWrapper) Warnf(format string, args ...interface{}) {
	logger.lw.Warnf(format, args)
}
func (logger *loggerWrapper) Debug(args ...interface{}) {
	logger.lw.Debug(args...)
}
func (logger *loggerWrapper) Debugf(format string, args ...interface{}) {
	logger.lw.Debugf(format, args)
}
func (logger *loggerWrapper) Printf(format string, args ...interface{}) {
	logger.lw.Infof(format, args)
}
func (logger *loggerWrapper) Println(args ...interface{}) {
	logger.lw.Info(args, "\n")
}