package sure_pkg_gen

import "github.com/yyle88/sure"

type SurePackageGenConfig struct {
	SourceRoot           string                 // SourceRoot is the root directory of the source code. // 源代码的根目录路径
	ErrorHandlingMode    sure.ErrorHandlingMode // ErrorHandlingMode specifies the error handling mode, such as strict (Must) or loose (Soft). // 指定错误处理模式，通常是严格（Must）或宽松（Soft）
	SourcePackagePath    string                 // SourcePackagePath is the import path of the source package. // 源代码包的导入路径
	ErrorHandlingPkgPath string                 // ErrorHandlingPkgPath is the path of the error handling package, default is "github.com/yyle88/sure". // 错误处理包的路径，默认为 "github.com/yyle88/sure"
	HandlerFuncReference string                 // HandlerFuncReference is the reference to the user-defined error handling functions like Soft and Must. // 用户自定义错误处理函数的引用，例如 Soft 和 Must
	NewPkgName           string                 // NewPkgName is the name of the new package. If not set, it will be generated from the source package name. // 新包名，如果未指定，将根据源代码包名生成
	OutputRoot           string                 // OutputRoot is the root directory where the generated code will be saved. // 生成代码保存的根目录
}

func NewSurePackageConfig(sourceRoot string, errorHandlingMode sure.ErrorHandlingMode, sourcePackagePath string) *SurePackageGenConfig {
	return &SurePackageGenConfig{
		SourceRoot:           sourceRoot,
		ErrorHandlingMode:    errorHandlingMode,
		SourcePackagePath:    sourcePackagePath,
		ErrorHandlingPkgPath: sure.GetPkgPath(), //默认用这个包 "github.com/yyle88/sure"
		HandlerFuncReference: sure.GetPkgName(), //默认使用 "sure" 调用软硬函数，比如 sure.Must(err) 和 sure.Soft(err) 因此很明显假如你有自己实现Must和Soft的话也可以用自己的
		NewPkgName:           "",                //默认不配置就会根据源码的包名拼接出新包名
		OutputRoot:           sourceRoot,
	}
}

func (c *SurePackageGenConfig) WithSourceRoot(sourceRoot string) *SurePackageGenConfig {
	c.SourceRoot = sourceRoot
	return c
}

func (c *SurePackageGenConfig) WithErrorHandlingMode(errorHandlingMode sure.ErrorHandlingMode) *SurePackageGenConfig {
	c.ErrorHandlingMode = errorHandlingMode
	return c
}

func (c *SurePackageGenConfig) WithSourcePackagePath(sourcePackagePath string) *SurePackageGenConfig {
	c.SourcePackagePath = sourcePackagePath
	return c
}

func (c *SurePackageGenConfig) WithErrorHandlingPkgPath(errorHandlingPkgPath string) *SurePackageGenConfig {
	c.ErrorHandlingPkgPath = errorHandlingPkgPath
	return c
}

func (c *SurePackageGenConfig) WithHandlerFuncReference(handlerFuncReference string) *SurePackageGenConfig {
	c.HandlerFuncReference = handlerFuncReference
	return c
}

func (c *SurePackageGenConfig) WithNewPkgName(newPkgName string) *SurePackageGenConfig {
	c.NewPkgName = newPkgName
	return c
}

func (c *SurePackageGenConfig) WithOutputRoot(outputRoot string) *SurePackageGenConfig {
	c.OutputRoot = outputRoot
	return c
}
