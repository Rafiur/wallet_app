package logger

import (
	"fmt"
	"github.com/Rafiur/wallet_app/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Logger methods interface
type Logger interface {
	InitLogger()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Panic(args ...interface{})
	Panicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	GetZapLogger() *zap.Logger
}

// apiLogger
type apiLogger struct {
	cfg    *config.Config
	logger *zap.Logger
}

// App Logger constructor
func NewApiLogger(cfg *config.Config) *apiLogger {
	return &apiLogger{cfg: cfg}
}

// For mapping config logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *apiLogger) getLoggerLevel(cfg *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// Init logger
func (l *apiLogger) InitLogger() {
	logLevel := l.getLoggerLevel(l.cfg)

	logWriter := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig
	if l.cfg.Logger.Development {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if l.cfg.Logger.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.logger = logger
}

// Logger methods
func (l *apiLogger) GetZapLogger() *zap.Logger {
	return l.logger
}

func (l *apiLogger) Debug(args ...interface{}) {
	l.logger.Debug("", argsToFields(args)...)
}

func (l *apiLogger) Debugf(template string, args ...interface{}) {
	l.logger.Sugar().Debugf(template, args...)
}

func (l *apiLogger) Info(args ...interface{}) {
	l.logger.Info("", argsToFields(args)...)
}

func (l *apiLogger) Infof(template string, args ...interface{}) {
	l.logger.Sugar().Infof(template, args...)
}

func (l *apiLogger) Warn(args ...interface{}) {
	l.logger.Warn("", argsToFields(args)...)
}

func (l *apiLogger) Warnf(template string, args ...interface{}) {
	l.logger.Sugar().Warnf(template, args...)
}

func (l *apiLogger) Error(args ...interface{}) {
	l.logger.Error("", argsToFields(args)...)
}

func (l *apiLogger) Errorf(template string, args ...interface{}) {
	//l.logger.Sugar().Errorf(template, args...)
	fmt.Println(template, args)
}

func (l *apiLogger) DPanic(args ...interface{}) {
	l.logger.DPanic("", argsToFields(args)...)
}

func (l *apiLogger) DPanicf(template string, args ...interface{}) {
	l.logger.Sugar().DPanicf(template, args...)
}

func (l *apiLogger) Panic(args ...interface{}) {
	l.logger.Panic("", argsToFields(args)...)
}

func (l *apiLogger) Panicf(template string, args ...interface{}) {
	l.logger.Sugar().Panicf(template, args...)
}

func (l *apiLogger) Fatal(args ...interface{}) {
	l.logger.Fatal("", argsToFields(args)...)
}

func (l *apiLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Sugar().Fatalf(template, args...)
}

func argsToFields(args ...interface{}) []zap.Field {
	fields := make([]zap.Field, len(args))
	for i, arg := range args {
		fields[i] = zap.Any(fmt.Sprintf("arg%d", i), arg)
	}
	return fields
}
