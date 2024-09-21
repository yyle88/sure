package sure_pkg_gen

import "github.com/yyle88/sure"

type Config struct {
	PkgRoot     string
	GenRoot     string
	SureEnum    sure.SureEnum
	PkgPath     string
	SurePkgPath string
	SureUseNode string
	NewPkgName  string
}

func NewConfig(pkgRoot string, sureEnum sure.SureEnum, pkgPath string) *Config {
	return &Config{
		PkgRoot:     pkgRoot,
		GenRoot:     pkgRoot,
		SureEnum:    sureEnum,
		PkgPath:     pkgPath,
		SurePkgPath: sure.GetPkgPath(), //默认用这个包 "github.com/yyle88/sure"
		SureUseNode: sure.GetPkgName(), //默认使用 "sure" 调用软硬函数，比如 sure.Must(err) 和 sure.Soft(err) 因此很明显假如你有自己实现Must和Soft的话也可以用自己的
		NewPkgName:  "",                //默认不配置就会根据源码的包名拼接出新包名
	}
}

func (c *Config) SetPkgRoot(pkgRoot string) *Config {
	c.PkgRoot = pkgRoot
	return c
}

func (c *Config) SetGenRoot(genRoot string) *Config {
	c.GenRoot = genRoot
	return c
}

func (c *Config) SetSureEnum(sureEnum sure.SureEnum) *Config {
	c.SureEnum = sureEnum
	return c
}

func (c *Config) SetPkgPath(pkgPath string) *Config {
	c.PkgPath = pkgPath
	return c
}

func (c *Config) SetSurePkgPath(surePkgPath string) *Config {
	c.SurePkgPath = surePkgPath
	return c
}

func (c *Config) SetSureUseNode(sureUseNode string) *Config {
	c.SureUseNode = sureUseNode
	return c
}

func (c *Config) SetNewPkgName(newPkgName string) *Config {
	c.NewPkgName = newPkgName
	return c
}
