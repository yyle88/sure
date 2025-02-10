package sure

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type ErrorHandlingMode string //意思是"柔性的"、"灵活的"，是遇到错误是崩溃，还是仅仅发出告警

//goland:noinspection GoSnakeCaseUsage
const (
	MUST ErrorHandlingMode = "Must" //硬硬的，出错时就崩溃
	SOFT ErrorHandlingMode = "Soft" //软软的，出错时仅告警
	OMIT ErrorHandlingMode = "Omit" //忽略的，出错时无视它
)

// Must 硬硬的，当有err时直接panic崩溃掉，流程中止
func Must(err error) {
	if err != nil {
		zaplog.LOGS.Skip1.Panic("must", zap.Error(err))
	}
}

// Soft 软软的，当有err时只打印个告警日志，流程继续
func Soft(err error) {
	if err != nil {
		zaplog.LOGS.Skip1.Warn("soft", zap.Error(err))
	}
}

// Omit 忽略的，当有err时不做任何提示动作，流程继续
func Omit(err error) {
	if err != nil {
		_ = err // 仅忽略错误
	}
}
