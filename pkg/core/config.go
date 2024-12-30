package core

import (
	"strings"

	"github.com/spf13/viper"
)

// OnInitialize 设置需要读取的配置文件名、环境变量，并将其内容读取到 viper 中.
func OnInitialize(configFile *string, envPrefix string, loadDirs []string, defaultConfigName string) func() {
	return func() {
		if configFile != nil {
			// 从命令行选项指定的配置文件中读取
			viper.SetConfigFile(*configFile)
		} else {
			for _, dir := range loadDirs {
				// 将 dir 目录加入到配置文件的搜索路径
				viper.AddConfigPath(dir)
			}

			// 设置配置文件格式为 YAML
			viper.SetConfigType("yaml")

			// 配置文件名称（没有文件扩展名）
			viper.SetConfigName(defaultConfigName)
		}

		// 读取匹配的环境变量
		viper.AutomaticEnv()

		// 设置环境变量前缀为 MINIBLOG
		viper.SetEnvPrefix(envPrefix)

		// 将 key 字符串中 '.' 和 '-' 替换为 '_'
		replacer := strings.NewReplacer(".", "_", "-", "_")
		viper.SetEnvKeyReplacer(replacer)

		// 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
		_ = viper.ReadInConfig()
	}
}
