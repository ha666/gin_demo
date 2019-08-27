package config

import (
	"github.com/ha666/logs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// InitConfig 读取并解析配置到Config对象
// 没有检查配置是否完整
func InitConfig() {

	// 启动时解析出错，程序退出
	err := parseConfig()
	if err != nil {
		logs.Emergency("解析配置文件出错:%s", err.Error())
	}

}

// parseConfig 解析并解密配置
func parseConfig() error {
	Config = new(config)
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, Config)
	if err != nil {
		return err
	}
	return nil
}
