# sure: 给现有 Go 代码添加断言和崩溃处理

`sure` 通过为现有的 Go 代码添加断言和崩溃处理功能来增强代码。它自动断言条件并在发生错误时崩溃，从而帮助你在不需要手动添加重复检查的情况下改善遗留代码的错误处理。

## 英文文档

[English README](README.md)

## 创作背景

[CREATION_IDEAS](internal/docs/CREATION_IDEAS.zh.md)

## 模块概述

### `sure_cls_gen`: **生成带有断言的 Go 类**

从预定义的对象生成 Go 类，再嵌入断言逻辑以防止常见错误。

### `sure_pkg_gen`: **生成带有错误处理的 Go 包**

从现有代码中提取函数并生成 Go 包，同时集成断言和崩溃处理。

### `cls_stub_gen`: **生成带有断言的 Go 方法存根**

给 Go 对象生成方法存根，再嵌入断言以确保适当的错误处理。

## 使用示例

### 示例：

- [sure_cls_gen](internal/examples/example_sure_cls_gen)
- [sure_pkg_gen](internal/examples/example_sure_pkg_gen)
- [cls_stub_gen](internal/examples/example_cls_stub_gen)

---

## 许可

`sure` 是一个开源项目，发布于 MIT 许可证下。有关更多信息，请参阅 LICENSE 文件。

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
