/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: yaml.go
 * @Date 2018/05/14
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package yaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Mysql db
type Mysql struct {
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Debug    bool   `yaml:"debug"`
}

// Redis db
type Redis struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Pass    string `yaml:"pass"`
	MaxConn int    `yaml:"maxConn"`
}

// Config define common yaml config
type Config struct {
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

// Init config from yaml file, conf can defined by yourself
func Init(file string, conf interface{}) error {
	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fileContent, conf)
	if err != nil {
		return err
	}

	return nil
}
