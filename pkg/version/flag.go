package version

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/pflag"
)

// versionValue 用于定义版本标识的类型.
type versionValue int

// 定义一些常量.
const (
	// 未设置版本.
	VersionNotSet versionValue = 0
	// 启用版本.
	VersionEnabled versionValue = 1
	// 原始版本.
	VersionRaw versionValue = 2
)

const (
	// 表示原始版本的字符串.
	strRawVersion = "raw"
	// 版本标志的名称.
	versionFlagName = "version"
)

// versionFlag 定义了版本标志.
var versionFlag = Version(versionFlagName, VersionNotSet, "Print version information and quit.")

func (v *versionValue) IsBoolFlag() bool {
	return true
}

func (v *versionValue) Get() interface{} {
	return v
}

// String 实现了 pflag.Value 接口中的 String 方法.
func (v *versionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion // 返回原始版本字符串
	}
	return strconv.FormatBool(bool(*v == VersionEnabled))
}

// Set 实现了 pflag.Value 接口中的 Set 方法.
func (v *versionValue) Set(s string) error {
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}
	boolVal, err := strconv.ParseBool(s)
	if boolVal {
		*v = VersionEnabled
	} else {
		*v = VersionNotSet
	}
	return err
}

// Type 实现了 pflag.Value 接口中的 Type 方法.
func (v *versionValue) Type() string {
	return "version"
}

// VersionVar 定义了一个具有指定名称和用法的标志.
func VersionVar(p *versionValue, name string, value versionValue, usage string) {
	*p = value
	pflag.Var(p, name, usage)

	// `--version` 等价于 `--version=true`
	pflag.Lookup(name).NoOptDefVal = "true"
}

// Version 包装了 VersionVar 函数.
func Version(name string, value versionValue, usage string) *versionValue {
	p := new(versionValue)
	VersionVar(p, name, value, usage)
	return p
}

// AddFlags 在任意 FlagSet 上注册这个包的标志，这样它们指向与全局标志相同的值.
func AddFlags(fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(versionFlagName))
}

// PrintAndExitIfRequested 将检查是否传递了 `--version` 标志，如果是，则打印版本并退出.
func PrintAndExitIfRequested() {
	// 检查版本标志的值并打印相应的信息
	if *versionFlag == VersionRaw {
		fmt.Printf("%s\n", Get().Text())
		os.Exit(0)
	} else if *versionFlag == VersionEnabled {
		fmt.Printf("%s\n", Get().String())
		os.Exit(0)
	}
}
