# mustdone
在我们开发golang代码时，经常会遇到比如 res, err := a.Run() 的情况，这时假如使用 res := amust.Run() 或者 res := a.Must().Run()岂不是能够避免频繁的判断 if err != nil 啦，这个包的目的就是提供这样的便利，当然本整活大师开发的 `github.com/yyle88/done` 也能解决问题，但毕竟不是还得多一层`nice`调用嘛，而这个工具将让代码自己提供错误时panic/ignore的选项，当然包名的话在mustsoft和mustgo和flexible间选择半天，最终想到也可以和`github.com/yyle88/done`套套近乎干脆就叫`mustdone`吧
