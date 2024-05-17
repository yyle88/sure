package mustdone

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type FlexibleHandlingType string

//goland:noinspection GoSnakeCaseUsage
const (
	MUST FlexibleHandlingType = "Must"
	SOFT FlexibleHandlingType = "Soft"
)

// Must 硬硬的，当有err时直接panic崩溃掉，流程中止
func Must(err error) {
	if err != nil {
		zaplog.LOG.Panic("must", zap.Error(err))
	}
}

// Soft 软软的，当有err时只打印个告警日志，流程继续
func Soft(err error) {
	if err != nil {
		zaplog.LOG.Warn("soft", zap.Error(err))
	}
}
