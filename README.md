# sure
在我们开发golang代码时，经常会遇到需要判断 err 非空的情况，但这有点麻烦，因此我发明了这个包，能够在确定err不会触发时直接碾过错误。

比如在
```
res, err := a.Run()
if err != nil {
    panic(err)
}
```
这个场景里，假如主逻辑不能运行，panic 有助于定位问题，让服务快速崩溃有助于外部检测，即时重启。

或者在
```
cfg, err := config.LoadFromFile(path)
if err != nil {
    panic(err)
}
```
这个场景里，假如读取配置报错，则系统已无法靠自身逻辑恢复，直接 panic，以便于运维同事发现问题。

这时假如使用
```
res := a_must.Run() //假设 a 是个包，而通过 a 包能得到 a_must 包，在里面自带出错时 panic 的函数。
```
或者
```
res := a.Must().Run() //假设 a 是个对象，能通过 A 类得到 AMust 类，里面自带出错时 panic 的方法。
``` 
就能避免频繁的判断 if err != nil 让程序变得更丝滑。

这种丝滑是指可以让代码维持链式调用。

比如原本的：
```
res, err := opt.GetR()
if err != nil {
    panic(err)
}
abc, err := res.GetA()
if err != nil {
    panic(err)
}
xyz, err := abc.GetX()
if err != nil {
    panic(err)
}
```
就可以写成这样的语句:
```
xyz := opt.Must().GetR().Must().GetA().Must().GetX()
```
这就比每次调用完判断是否有 error 简单些，在略微非正式的情况下是无妨的。

这个包的目的就是提供这样的便利。

# 提供类和包两种情况下的 sure 操作
有的方法是某个类的成员方法，比如 `param.Check()`，当参数不正确时报错，当联调结束以后参数基本都是对的，即使出错 panic 也没问题。

而有的函数是某个包的小函数，比如 `json.Marshal` 函数，它就几乎不会出错(除非传个接口给它)，经常需要判断err是否非空，其实没必要。

因此对于类和包，两种情况我做了两个生成器。

## 类操作代码生成器
假设我们封装了个类 A 它有:
```
GetConfig(path string) (Config, error)
```
就简单地封装这个操作为这样:
```
cfg := a.Must().GetConfig(cfgPath)
```
这样岂不是非常方便，这就是“类操作生成器”的基本逻辑，就是把类中所有导出方法都在遇到 err 时 panic，在调用时就能省去判断逻辑。

详情见 demos:

[Demo1](/internal/examples/example1)

[Demo4](/internal/examples/example4)

[Demo5](/internal/examples/example5)

## 包操作代码生成器
假如封装的函数在 `utils` 包里，常规的调用是这样的:
```
cfg, err := utils.GetConfig(cfgPath)
```
经过代码生成以后会得到 `utils_must` 新包，调用就被简化为这样:
```
cfg := utils_must.GetConfig(cfgPath)
```

详情见 demos:

[Demo2](/internal/examples/example2)

[Demo3](/internal/examples/example3)

## 思路
[创作背景](/internal/docs/CREATION_IDEAS.md)

## 最终:
Give me stars. Thank you!!!
