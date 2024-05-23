# 创作背景-我想过的其它方案

## 设计意图
这个工具将让代码自己提供错误时panic/warning的选项。

## 基本函数
该包的核心就是提供两个函数:
```
// Must 硬硬的，当有err时直接panic崩溃掉，流程中止
func Must(err error)

// Soft 软软的，当有err时只打印个告警日志，流程继续
func Soft(err error)
```

## 简单调用
在代码逻辑里就可以这样使用:
```
data, err := json.Marshal(example)
mustdone.Must(err) //当没有错误时就什么也不做，当出错时将 panic 崩溃
```
或者使用:
```
mustdone.Soft(err) //当没有错误时就什么也不做，当出错时将会打印 waring 日志
```
当然很明显的，像这样粗暴的抛出异常，或者直接忽略错误，在正式场景下可能是不适用的，需要注意。

## 代码生成
使用代码生成自动赋予某类中函数 `must` 的能力，或者某个包中函数 `must` 的能力。

## 其它背景：
在开发时也遇到了个非常困惑的事情，就是，我见别人定义的Must函数往往是这样使用的 `a := Must(defaultValue)` 即假如计算错误就返回默认值，即哪怕没有结果也必须返回结果。

比如这个函数：
```
func (j *Json) MustInt(args ...int) int {
	var def int

	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	default:
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}

	i, err := j.Int()
	if err == nil {
		return i
	}

	return def
}
```
我当时看了很久才理解这里面 `must` 的意思，他甚至把默认值做成可选的，看来这里的 `must` 跟我理解的不同。

而真正需要在出错时直接告警的，他们用的是 Require 函数，我滴神啊，看来英文不好的人是连 must 是啥意思都不懂，也是够悲催的。

我认为的 `must` 是，你必须做成某件事得到结果，否则就要接受惩罚（`panic`），而他们理解的 `must` 是，我不管你遇到什么困难你都要给结果，假如得不到结果就返回给你的默认值，他们更关注是必然得到结果。

既然对 `must` 的理解有歧义的话这里我就说明白，我的 `must` 就是必须做成某件事，而且没有遇到错误情况，否则就会 `panic`，整个项目都是基于这个语境做的。

比如：
```
func (T *SimpleMust) Strings(key string) (res []string) {
	res, err1 := T.T.Strings(key)
	mustdone.Must(err1)
	return res
}
```
调用：
```
tags := sim.Must().Strings("tags")
```
假如得不到就直接 `panic` 崩溃，相当于在获取值的时候，附带断言的效果。

整个项目都是基于这个语境做的。

## 相似的包
当然包名的话在mustsoft和mustgo和flexible间选择半天，最终想到也可以和`github.com/yyle88/done`套套近乎干脆就叫`mustdone`吧

当然本整活大师开发的 `github.com/yyle88/done` 也能解决问题，但毕竟不是还得多一层`nice`调用嘛，而且每次都在外面包一层，当包装调用超过两层时就会让代码变得不易读。

比如：
```
defer func() { //这是关闭 db *gorm.DB 时的操作，假设时在测试用例里，起了个临时DB，结束时就可以这样关闭它
    done.Done(done.VCE(db.DB()).Nice().Close())
}()
```
这段代码的可读性就已经大大下降啦。

## 最终效果
其实，使用场景时很少的，而且在外部重新生成代码赋能的方法在 90% 的场景下都是不如 `github.com/yyle88/done` 的，而剩下 10% 的场景直接用 `if err != nil { panic(err) }` 就行。

当然假如特别想用也可以用用，毕竟关键时刻能省三行代码也不错。
