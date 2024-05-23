# mustdone
在我们开发golang代码时，经常会遇到
```
res, err := a.Run()
if err != nil {
    panic(err)
}
```
的场景。比如这种场景：在main函数里，需要读个配置，假如读取出错就直接崩溃，这种操作就是有利于发现问题的，也有利于代码的简洁。

这时假如使用
```
res := a_must.Run() //假设 a 是个包，而通过 a 包能得到 a_must 包，里面自带出错时 panic 的逻辑。
```
或者
```
res := a.Must().Run() //假设 a 是个对象，能通过 A 类得到 AMust 类，里面自带出错时 panic 的逻辑。
``` 
岂不是能够避免频繁的判断 if err != nil 让程序变得更丝滑。

这种丝滑是指可以让代码维持链式调用。

比如原本的：
```
res, err := opt.GetRes()
if err != nil {
    panic(err)
}
abc, err := res.GetAbc()
if err != nil {
    panic(err)
}
xyz, err := abc.GetXyz()
if err != nil {
    panic(err)
}
```
就可以这样:
```
xyz := opt.Must().GetRes().Must().GetAbc().Must().GetXyz()
```
这就比每次调用完判断是否有 error 简单些。

这个包的目的就是提供这样的便利。

## 代码生成（类操作）
假设我们封装了个类 A 它有:
```
GetConfig(path string) (Config, error)
```
而我们调用的时候常常这样用:
```
cfg, err := a.GetConfig(cfgPath)
mustdone.Must(err) //读不到配置就直接退出
```
就简单地封装这个操作为这样:
```
cfg := a.Must().GetConfig(cfgPath)
```
这样岂不是非常方便，这就是“类操作生成逻辑”。

[Demo1](/internal/examples/example1)
[Demo4](/internal/examples/example4)
[Demo5](/internal/examples/example5)

## 代码生成（包操作）
同样的，假如封装的函数在 `utils` 包里，常规的调用是这样的:
```
cfg, err := utils.GetConfig(cfgPath)
mustdone.Must(err) //读不到配置就直接退出
```
经过代码生成以后会得到 `utils_must` 新包，调用就被简化为这样:
```
cfg := utils_must.GetConfig(cfgPath)
```
[Demo2](/internal/examples/example2)
[Demo3](/internal/examples/example3)

## 思路
[创作背景](/internal/docs/CREATION_IDEAS.md)

## 最终:
Give me stars. Thank you!!!
