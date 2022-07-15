package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"moul.io/zapfilter"
)

func main() {
	//c := zap.NewExample().Core()
	//
	//logger := zap.New(zapfilter.NewFilteringCore(c, zapfilter.MustParseRules("demo*")))
	//defer logger.Sync()
	//
	//logger.Debug("hello world!")
	//logger.Named("demo").Debug("hello earth!")
	//logger.Named("other").Debug("hello universe!")

	filtered := zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapfilter.NewFilteringCore(c, zapfilter.MustParseRules("demo*"))
	})
	logger, _ := zap.NewDevelopment(filtered)
	defer logger.Sync()
	//logger.WithOptions(filtered).Debug("hello world!")
	//logger.WithOptions(filtered).Named("demo").Debug("hello earth!")
	//logger.WithOptions(filtered).Named("other").Debug("hello universe!")
	zap.ReplaceGlobals(logger)
	zap.S().Named("demo").Infof("demo oooo")
}
