package listener

import (
	"Notifier/helpers"
	"gopkg.in/ini.v1"
	"os"
)

type Config struct {
	host     string
	port     string
	dbName   string
	username string
	password string
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
		host:     getValueWithDefault(section, "host", "localhost"),
		port:     getValueWithDefault(section, "port", "5432"),
		username: getValueWithDefault(section, "username", "admin"),
		password: getValueWithDefault(section, "password", "admin"),
		dbName:   getValue(section, "dbname"),
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
