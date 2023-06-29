package mocks

import "time"

type ConfigMock struct{}

func (config *ConfigMock) GetServerURL() string {
	return "localhost:4000"
}

func (config *ConfigMock) GetAddress() string {
	return "localhost"
}

func (config *ConfigMock) GetPort() string {
	return "4000"
}

func (config *ConfigMock) GetServerShutdownTimeout() time.Duration {
	return 30 * time.Second
}

func (config *ConfigMock) GetCBRWSDLTimeout() time.Duration {
	return 5 * time.Second
}

func (config *ConfigMock) GetCBRWSDLAddress() string {
	return ""
}

func (config *ConfigMock) GetLoggingOn() bool {
	return true
}

func (config *ConfigMock) GetDateTimeResponseLayout() string {
	return "2006-01-02"
}

func (config *ConfigMock) GetDateTimeRequestLayout() string {
	return "2006-01-02"
}

func (config *ConfigMock) GetPermittedRequests() map[struct{}]string {
	return nil
}
