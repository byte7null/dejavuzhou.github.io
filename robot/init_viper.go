package robot

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	
	viper.AutomaticEnv() // read in environment variables that match
	
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("加载配置文件失败:", viper.ConfigFileUsed(), err)
	}
	initLogrus()
}
