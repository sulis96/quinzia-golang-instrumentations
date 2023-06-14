package config

import "os"

type (
	AppConfig struct {
		AppName        string
		HTTPServerPort string
	}

	DbConfig struct {
		Host     string
		Port     string
		Username string
		Password string
		DbName   string
		Schema   string
		Driver   string
	}

	InstrumentationConfig struct {
		ServiceName              string
		InsecureMode             string
		OtelExporterOtlpEndpoint string
	}
)

func InitAppConfig() *AppConfig {
	return &AppConfig{
		AppName:        os.Getenv("APP_NAME"),
		HTTPServerPort: os.Getenv("HTTP_SERVER_PORT"),
	}
}

func InitDatabaseConfig() *DbConfig {
	return &DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		Schema:   os.Getenv("DB_SCHEMA"),
		Driver:   os.Getenv("DB_DRIVER"),
	}
}

func InitInstrumentationConfig() *InstrumentationConfig {
	return &InstrumentationConfig{
		ServiceName:              os.Getenv("SERVICE_NAME"),
		InsecureMode:             os.Getenv("INSECURE_MODE"),
		OtelExporterOtlpEndpoint: os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"),
	}
}
