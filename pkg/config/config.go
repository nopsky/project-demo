/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/9/16 19:19
 */
package config

import (
	"bytes"
	"io/ioutil"

	"github.com/nopsky/project-demo/pkg/logger"
	"github.com/spf13/viper"
)

var conf *viper.Viper

func Load(path string) error {
	var (
		err error
		v   = viper.New()
	)

	v.AddConfigPath(".")
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	logger.Infof("use config file -> %s\n", v.ConfigFileUsed())

	includes := v.GetStringSlice("includes")
	for _, i := range includes {

		fd, err := ioutil.ReadFile(i)
		if err != nil {
			logger.Error("load config err:", err.Error())
			return err
		}

		v.MergeConfig(bytes.NewReader(fd))

		logger.Infof("load config file -> %s\n", i)

	}

	conf = v

	return err
}

func UnmarshalKey(key string, val interface{}) error {
	return conf.UnmarshalKey(key, val)
}
