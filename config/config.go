//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

// ConfigurationModel represent the configuration model
type ConfigurationModel struct {
	Port    string `json:"port"`
	MongoDB struct {
		Addr     string `json:"addr"`
		Database string `json:"database"`
	} `json:"mongo_db"`
}

func InitConfig() (*ConfigurationModel, error) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	basepath = strings.Replace(basepath, "config", "", -1)
	file := basepath + "config.json"
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Failed to load auth configuration file: %s", err.Error())
	}

	var configuration *ConfigurationModel
	err = json.Unmarshal(raw, configuration)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse auth configuration file: %s", err.Error())
	}

	return configuration, nil
}
