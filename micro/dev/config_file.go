package main

import (
	"apm/pkg/util"
	"encoding/json"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/logger"
	"strings"
)

type fileConfig struct {
	config map[string]json.RawMessage
}

// NewConfig returns new config
func NewConfig(fileName string) (*fileConfig, error) {
	config := make(map[string]json.RawMessage)

	err := util.ReadFile(&config, fileName)
	if err != nil {
		return nil, err
	}
	logger.Debug("config => ", config)
	return &fileConfig{
		config: config,
	}, nil
}

func formatKey(v string) string {
	if len(v) == 0 {
		return ""
	}

	v = strings.ToLower(v)
	return strings.Replace(v, ".", "_", -1)
}

func (c *fileConfig) Get(path string, options ...config.Option) (config.Value, error) {
	name := formatKey(path)
	v, exit := c.config[name]
	if !exit {
		v = json.RawMessage{}
	}
	data, err := v.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return config.NewJSONValue(data), nil
}

func (c *fileConfig) Set(path string, val interface{}, options ...config.Option) error {
	key := formatKey(path)
	raw, err := json.Marshal(val)
	if err != nil {
		return err
	}
	c.config[key] = json.RawMessage(raw)
	return nil
}

func (c *fileConfig) Delete(path string, options ...config.Option) error {
	v := formatKey(path)
	delete(c.config, v)
	return nil
}
