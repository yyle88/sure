package sure

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type SureEnum string //意思是"柔性的"、"灵活的"，是遇到错误是崩溃，还是仅仅发出告警

//goland:noinspection GoSnakeCaseUsage
const (
	MUST SureEnum = "Must" //硬硬的，出错时就崩溃
	SOFT SureEnum = "Soft" //软软的，出错时仅告警
	OMIT SureEnum = "Omit" //忽略的，出错时无视它
)

// Must 硬硬的，当有err时直接panic崩溃掉，流程中止
func Must(err error) {
	if err != nil {
		zaplog.LOGS.P1.Panic("must", zap.Error(err))
	}
}

// Soft 软软的，当有err时只打印个告警日志，流程继续
func Soft(err error) {
	if err != nil {
		zaplog.LOGS.P1.Warn("soft", zap.Error(err))
	}
}

// Omit 忽略的，当有err时不做任何提示动作，流程继续
func Omit(err error) {
	if err != nil {
		_ = err // 仅忽略错误
	}
}
