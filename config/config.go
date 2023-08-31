/*
 * *******************************************************************
 * @项目名称: config
 * @文件名称: config.go
 * @Date: 2020/02/07
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package config

import (
	"os"
	"strings"

	"git.bhex.io/bhpc/wallet/common/logger"
	"git.bhex.io/bhpc/wallet/common/yaml"
)

const (
	instanceName = "INSTANCE_NAME"
	appName      = "baasnode-server"
	projectName  = "baasnode"
	prodFile     = "conf/config_prod.yml"
	testFile     = "conf/config_test.yml"
)

var (
	// Config global config object
	Config baasConfig
	log    = logger.New("config")
)

// baasConfig define vds yaml
type baasConfig struct {
	Redis struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Pass    string `yaml:"pass"`
		MaxConn int    `yaml:"maxConn"`
	} `yaml:"redis"`
	Chainnode struct {
		Protocol string `yaml:"protocol"`
		Port     string `yaml:"port"`
		HTTPPort string `yaml:"httpport"`
	} `yaml:"baasnode"`
	Wallet struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"wallet"`
	Risk struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"risk"`
}

func init() {
	var configPath, configFile string

	if os.Getenv(instanceName) == appName {
		configFile = prodFile
	} else {
		configFile = testFile
	}

	if !fileExist(configFile) {
		// get current dir
		currentDir, _ := os.Getwd()

		// get actual config path
		index := strings.Index(currentDir, projectName)
		configPath = string(currentDir[:index+len(projectName)]) + "/" + configFile
	} else {
		configPath = configFile
	}

	// read config
	err := yaml.Init(configPath, &Config)
	if err != nil {
		log.Errorf("[ALERT] Init config err:%v", err)
		panic(err)
	}
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
