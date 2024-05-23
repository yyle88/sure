package example5mustsoftnode

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

var NODE = &Node{}

type Node struct{}

// Must 硬硬的，当有err时直接panic崩溃掉，流程中止
func (node *Node) Must(err error) {
	if err != nil {
		zaplog.LOGS.P1.Panic("must", zap.Error(err))
	}
}

// Soft 软软的，当有err时只打印个告警日志，流程继续
func (node *Node) Soft(err error) {
	if err != nil {
		zaplog.LOGS.P1.Warn("soft", zap.Error(err))
	}
}
