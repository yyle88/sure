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

给 Go 类型创建封装单例结构体方法的包级函数，简化访问使用。。

## 使用示例

### 示例：

- [使用 `sure_cls_gen` 生成类](internal/examples/example_sure_cls_gen)
- [使用 `sure_pkg_gen` 生成包](internal/examples/example_sure_pkg_gen)
- [使用 `cls_stub_gen` 生成单例](internal/examples/example_cls_stub_gen)

---

## 许可证类型

项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE)。

---

## 贡献新代码

非常欢迎贡献代码！贡献流程：

1. 在 GitHub 上 Fork 仓库 （通过网页界面操作）。
2. 克隆Forked项目 (`git clone https://github.com/yourname/repo-name.git`)。
3. 在克隆的项目里 (`cd repo-name`)
4. 创建功能分支（`git checkout -b feature/xxx`）。
5. 添加代码 (`git add .`)。
6. 提交更改（`git commit -m "添加功能 xxx"`）。
7. 推送分支（`git push origin feature/xxx`）。
8. 发起 Pull Request （通过网页界面操作）。

请确保测试通过并更新相关文档。

---

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
