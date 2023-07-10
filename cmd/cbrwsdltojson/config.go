package main

import (
	"errors"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	permittedRequest       map[string]struct{} `mapstructure:"PERMITTED_REQUESTS"`
	Logger                 LoggerConf          `mapstructure:"Logger"`
	ServerShutdownTimeout  time.Duration       `mapstructure:"SERVER_SHUTDOWN_TIMEOUT"`
	CBRWSDLTimeout         time.Duration       `mapstructure:"CBR_WSDL_TIMEOUT"`
	address                string              `mapstructure:"ADDRESS"`
	port                   string              `mapstructure:"PORT"`
	cbrWSDLAddress         string              `mapstructure:"CBR_WSDL_ADDRESS"`
	dateTimeResponseLayout string              `mapstructure:"DATE_TIME_RESPONSE_LAYOUT"`
	dateTimeRequestLayout  string              `mapstructure:"DATE_TIME_REQUEST_LAYOUT"`
	loggingOn              bool                `mapstructure:"LOGGING_ON"`
}

type LoggerConf struct {
	Level string `mapstructure:"LOG_LEVEL"`
}

func NewConfig() Config {
	return Config{}
}

func (config *Config) Init(path string) error {
	if path == "" {
		err := errors.New("void path to config.env")
		return err
	}

	viper.SetDefault("ADDRESS", "cbrwsdltojson")
	viper.SetDefault("PORT", "4000")
	viper.SetDefault("SERVER_SHUTDOWN_TIMEOUT", 30*time.Second)
	viper.SetDefault("CBR_WSDL_TIMEOUT", 5*time.Second)
	viper.SetDefault("LOGGING_ON", true)
	viper.SetDefault("CBR_WSDL_ADDRESS", "http://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx")
	viper.SetDefault("DATE_TIME_RESPONSE_LAYOUT", "2006-01-02 15:04:05")
	viper.SetDefault("DATE_TIME_REQUEST_LAYOUT", "2006-01-02 15:04:05")
	viper.SetDefault("PERMITTED_REQUESTS", "")

	viper.SetDefault("LOG_LEVEL", "debug")

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok { //nolint:errorlint
			return err
		}
	}
	config.address = viper.GetString("ADDRESS")
	config.port = viper.GetString("PORT")
	config.ServerShutdownTimeout = viper.GetDuration("SERVER_SHUTDOWN_TIMEOUT")
	config.CBRWSDLTimeout = viper.GetDuration("CBR_WSDL_TIMEOUT")
	config.loggingOn = viper.GetBool("LOGGING_ON")
	config.cbrWSDLAddress = viper.GetString("CBR_WSDL_ADDRESS")
	config.dateTimeResponseLayout = viper.GetString("DATE_TIME_RESPONSE_LAYOUT")
	config.dateTimeRequestLayout = viper.GetString("DATE_TIME_REQUEST_LAYOUT")
	tempPermReq := viper.GetString("PERMITTED_REQUESTS")
	permittedRequests := make(map[string]struct{})
	if tempPermReq != "" {
		tempPermReqSl := strings.Split(tempPermReq, " ")
		for _, curPR := range tempPermReqSl {
			permittedRequests[strings.TrimSpace(curPR)] = struct{}{}
		}
	}
	config.permittedRequest = permittedRequests
	return nil
}

func (config *Config) GetServerURL() string {
	return config.address + ":" + config.port
}

func (config *Config) GetAddress() string {
	return config.address
}

func (config *Config) GetPort() string {
	return config.port
}

func (config *Config) GetServerShutdownTimeout() time.Duration {
	return config.ServerShutdownTimeout
}

func (config *Config) GetCBRWSDLTimeout() time.Duration {
	return config.CBRWSDLTimeout
}

func (config *Config) GetCBRWSDLAddress() string {
	return config.cbrWSDLAddress
}

func (config *Config) GetLoggingOn() bool {
	return config.loggingOn
}

func (config *Config) GetDateTimeResponseLayout() string {
	return config.dateTimeResponseLayout
}

func (config *Config) GetDateTimeRequestLayout() string {
	return config.dateTimeRequestLayout
}

func (config *Config) GetPermittedRequests() map[string]struct{} {
	return config.permittedRequest
}
