package listener

import (
	"Notifier/helpers"
	"gopkg.in/ini.v1"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Host         string
	Port         string
	DBName       string
	Username     string
	Password     string
	MinReconnect time.Duration
	MaxReconnect time.Duration
	TriggerName  string
	WorkerCount  int
}

func LoadConfig(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		helpers.RaiseStdErr("error while parsing config file :" + err.Error())
	}
	return unmarshalConfig(data)
}

func unmarshalConfig(data []byte) *Config {
	//todo: get config section
	cfg, err := ini.Load(data)
	helpers.CheckErr("error while parsing configs", err)
	if cfg.HasSection("postgresql") {
		// better error handling
		return dispatch(cfg.Section("postgresql"))
	}
	return &Config{}
}

func dispatch(section *ini.Section) *Config {
	return &Config{
		Host:     getValueWithDefault(section, "host", "localhost"),
		Port:     getValueWithDefault(section, "port", "5432"),
		Username: getValueWithDefault(section, "username", "admin"),
		Password: getValueWithDefault(section, "password", "admin"),
		DBName:   getValue(section, "dbname"),
		MinReconnect: getValueWithDefaultAndConvert(section, "minReconnect", "10", func(data string) interface{} {
			duration, err := strconv.Atoi(data)
			helpers.PureCheckErr(err)
			return time.Duration(duration) * time.Second
		}).(time.Duration),
		MaxReconnect: getValueWithDefaultAndConvert(section, "minReconnect", "60", func(data string) interface{} {
			duration, err := strconv.Atoi(data)
			helpers.PureCheckErr(err)
			return time.Duration(duration) * time.Second
		}).(time.Duration),
		TriggerName: getValue(section, "triggerName"),
		WorkerCount: getValueWithDefaultAndConvert(section, "workerCount", "5", func(data string) interface{} {
			count, err := strconv.Atoi(data)
			helpers.PureCheckErr(err)
			return count
		}).(int),
	}
}

func getValueWithDefaultAndConvert(section *ini.Section, key string, defaultValue string, converter func(data string) interface{}) interface{} {
	res := getValueWithDefault(section, key, defaultValue)
	return converter(res)

}

func getValueWithDefault(section *ini.Section, key string, defaultValue string) string {
	if section.HasKey(key) {
		return getValue(section, key)
	}
	return defaultValue
}

func getValue(section *ini.Section, key string) string {
	res, err := section.GetKey(key)
	helpers.CheckErr("error while getting key"+key+" from config file", err)
	return res.Value()
}
