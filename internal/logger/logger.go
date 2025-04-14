package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func InitLogger() {
	l, _ := zap.NewDevelopment()
	Log = l.Sugar()
}
