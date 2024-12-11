package LoggerWordbox

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Info(message string, args ...interface{}) {
	zap.L().Sugar().Infof(message, args...)
}

func Error(message string, args ...interface{}) {
	zap.L().Sugar().Errorf(message, args...)
}

func Warn(message string, args ...interface{}) {
	zap.L().Sugar().Warnf(message, args...)
}

func Debug(message string, args ...interface{}) {
	zap.L().Sugar().Debugf(message, args...)
}

func Panic(message string, args ...interface{}) {
	zap.L().Sugar().Panicf(message, args...)
}

func Fatal(message string, args ...interface{}) {
	zap.L().Sugar().Fatalf(message, args...)
}

func CreateLogger() {
	loggerCfg := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	zap.ReplaceGlobals(zap.Must(loggerCfg.Build(zap.AddStacktrace(zap.ErrorLevel))))
}

var encoderConfig = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "severity",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "message",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    encodeLevel(),
	EncodeTime:     zapcore.RFC3339TimeEncoder,
	EncodeDuration: zapcore.MillisDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func encodeLevel() zapcore.LevelEncoder {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			enc.AppendString("INFO")
		case zapcore.WarnLevel:
			enc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			enc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			enc.AppendString("CRITICAL")
		case zapcore.PanicLevel:
			enc.AppendString("ALERT")
		case zapcore.FatalLevel:
			enc.AppendString("EMERGENCY")
		}
	}
}
