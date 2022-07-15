package initialize

import (
	"go.uber.org/zap"
)

func InitLogger() {
	//filtered := zap.WrapCore(func(c zapcore.Core) zapcore.Core {
	//	return zapfilter.NewFilteringCore(c, zapfilter.MustParseRules("demo*"))
	//})
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
