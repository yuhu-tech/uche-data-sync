package config

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/yuhu-tech/uche-data-sync/internal/common/log"
	"github.com/yuhu-tech/uche-data-sync/internal/common/toml"
)

var configFile string

type Config struct {
	LogDir   string `toml:"log_dir"`
	LogLevel string `toml:"log_level"`

	Mqtt       mqtt       `toml:"mqtt"`
	MongoDB    mongodb    `toml:"mongo_db"`
	WorkerPool workerPool `toml:"worker_pool"`
}

type mqtt struct {
	Server                string   `toml:"server"`
	ClientID              string   `toml:"client_id"`
	Username              string   `toml:"username"`
	Password              string   `toml:"password"`
	Topic                 []string `toml:"topic"`
	QoS                   []byte   `toml:"qos"`
	DisconnectWaitingTime uint     `toml:"disconnct_waitng_time"`
}

type mongodb struct {
	URL              string `toml:"url"`
	DB               string `toml:"db"`
	MongoDBPoolLimit int    `toml:"pool_limit"`
}

type workerPool struct {
	WorkSize int `toml:"work_size"`
	PoolSize int `toml:"pool_size"`
}

func init() {
	flag.StringVar(&configFile, "conf", "", "Config file path")
	fmt.Println(configFile)
}

func InitConfig() (*Config, error) {
	var conf Config

	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		return nil, err
	}

	logPath := filepath.Join(conf.logDir(), "collection.log")

	log.SetFlags(log.Lshortfile)
	log.Init(logPath, conf.logLevel())
	log.Info("Collection configuration: %v", conf)
	return &conf, nil
}

func (c *Config) logDir() string {
	return c.LogDir
}

func (c *Config) logLevel() string {
	return c.LogLevel
}

func (c *Config) Server() string {
	return c.Mqtt.Server
}

func (c *Config) ClientID() string {
	return c.Mqtt.ClientID
}

func (c *Config) Username() string {
	return c.Mqtt.Username
}

func (c *Config) Password() string {
	return c.Mqtt.Password
}

func (c *Config) TopicandQoS() (map[string]byte, error) {
	topic := c.Mqtt.Topic
	qos := c.Mqtt.QoS

	if len(topic) != len(qos) || len(topic) == 0 {
		return nil, errors.New("Configuration error! Please check the config, then restart the application.")
	}

	filters := make(map[string]byte)
	for i := range topic {
		filters[topic[i]] = qos[i]
	}

	return filters, nil
}

func (c *Config) DisconnectWaitingTime() uint {
	return c.Mqtt.DisconnectWaitingTime
}

func (c *Config) URL() string {
	return c.MongoDB.URL
}

func (c *Config) DB() string {
	return c.MongoDB.DB
}

func (c *Config) MongoDBPoolLimit() int {
	return c.MongoDB.MongoDBPoolLimit
}

func (c *Config) WorkSize() int {
	return c.WorkerPool.WorkSize
}

func (c *Config) PoolSize() int {
	return c.WorkerPool.PoolSize
}
