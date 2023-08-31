/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: log.go
 * @Date: 2018/05/16
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package logger

import (
	"io/ioutil"
	"strings"

	"git.bhex.io/bhpc/wallet/common/env"

	"github.com/hhkbp2/go-logging"
	yaml "gopkg.in/yaml.v2"
)

const (
	defaultYAML = `root:
    level: DEBUG
    handlers: [h]
loggers:
    access_log:
        level: DEBUG
        handlers: [access_log]
        propagate: false

formatters:
    f:
        format: "%(asctime)s | | %(name)s:%(filename)s:%(lineno)d | %(levelname)s | %(funcname)s | %(message)s"
        datefmt: "%Y-%m-%d %H:%M:%S.%3n"
    f2:
        format: "%(asctime)s | | %(filename)s:%(lineno)d | %(levelname)s | %(message)s"
        datefmt: "%Y-%m-%d %H:%M:%S.%3n"

handlers:
    h:
        class: TimedRotatingFileHandler
        filepath: "logs/log.log"
        mode: O_APPEND
        bufferSize: 0
        when: "D"
        interval: 1
        utc: false
        backupCount: 10
        maxBytes: 524288000
        formatter: f
    access_log:
        class: TimedRotatingFileHandler
        filepath: "logs/access_log.log"
        mode: O_APPEND
        bufferSize: 0
        when: "D"
        interval: 1
        utc: false
        backupCount: 10
        maxBytes: 104857600
        formatter: f2
`
)

// Logger contains all logging method
type Logger interface {
	logging.Logger
}

func init() {
	var conf logging.Conf
	config, err := ioutil.ReadAll(strings.NewReader(defaultYAML))
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(config, &conf); err != nil {
		panic(err)
	}
	if err := logging.DictConfig(&conf); err != nil {
		panic(err)
	}

}

// New return a logger object with module name
func New(module string) Logger {
	logger := logging.GetLogger(module)

	if env.IsTestEnv() {
		logger.SetLevel(logging.LevelDebug)
	} else {
		logger.SetLevel(logging.LevelInfo)
	}

	return logger
}

// Close the logger object
func Close() {
	logging.Shutdown()
}
