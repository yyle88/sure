# mustdone
在我们开发golang代码时，经常会遇到比如 res, err := a.Run() 的情况。

这时假如使用 res := amust.Run() 或者 res := a.Must().Run()岂不是能够避免频繁的判断 if err != nil 啦。

这个包的目的就是提供这样的便利。

当然本整活大师开发的 `github.com/yyle88/done` 也能解决问题，但毕竟不是还得多一层`nice`调用嘛。

而这个工具将让代码自己提供错误时panic/ignore的选项。

当然包名的话在mustsoft和mustgo和flexible间选择半天，最终想到也可以和`github.com/yyle88/done`套套近乎干脆就叫`mustdone`吧

该包的核心就是提供两个函数
```
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
```
而且通过代码生成的逻辑让其它包也具备这种能力，即出错时要么崩溃要么告警，避免开发者反复处理各种error情况。

### 其它：
在开发时也遇到了个非常困惑的事情，就是，我见别人定义的Must函数往往是这样使用的 `a := Must(dft)` 即假如计算错误就返回默认值，即必须返回结果让外部使用。

而真正需要在出错时直接告警的，他们用的是 Require 函数，我滴神啊，看来英文不好的人是连 must 是啥意思都不懂，也是够悲催的。

我认为的 `must` 是，你必须做成某件事得到结果，否则就要接受惩罚（`panic`），而他们理解的 `must` 是，我不管你遇到什么困难你都要给结果，他们更关注结果。

既然对 `must` 的理解有歧义的话这里我就说明白，我的 `must` 就是必须做成某件事，而且没有遇到错误情况，否则就会 `panic`，整个项目都是基于这个语境做的。

### 最终:
Give me stars. Thank you!!!
