package mustdone

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type FlexibleEnum string //意思是"柔性的"、"灵活的"，是遇到错误是崩溃，还是仅仅发出告警

//goland:noinspection GoSnakeCaseUsage
const (
	MUST FlexibleEnum = "Must" //硬硬的，出错时就崩溃
	SOFT FlexibleEnum = "Soft" //软软的，出错时仅告警
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
