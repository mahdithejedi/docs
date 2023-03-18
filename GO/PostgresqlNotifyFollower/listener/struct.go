package listener

import (
	"Notifier/helpers"
	"gopkg.in/ini.v1"
	"os"
)

type Config struct {
	Host     string
	Port     string
	DBName   string
	Username string
	Password string
	//triggerName string
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
	}
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
