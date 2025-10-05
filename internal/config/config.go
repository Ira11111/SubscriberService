package config

import (
	"errors"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const cfgVar = "CONFIG_PATH"

type Config struct {
	Env        string       `yaml:"env" env-default:"dev"`
	DB         DBConfig     `yaml:"db" env-required:"true"`
	HttpServer ServerConfig `yaml:"server" env-required:"true"`
}

type DBConfig struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string
	Database string `yaml:"database"`
	SSLMode  string `yaml:"ssl_mode"`
}
type ServerConfig struct {
	Host            string           `yaml:"host" env-default:"localhost"`
	Port            int              `yaml:"port" env-default:"8080"`
	ReadTimeout     time.Duration    `yaml:"read_timeout" env-default:"30s"`
	WriteTimeout    time.Duration    `yaml:"write_timeout" env-default:"30s"`
	IdleTimeout     time.Duration    `yaml:"idle_timeout" env-default:"60s"`
	ShutdownTimeout time.Duration    `yaml:"shutdown_timeout" env-default:"10s"`
	Middleware      MiddlewareConfig `yaml:"middleware"`
}

type MiddlewareConfig struct {
	Logger   LoggerConfig   `yaml:"logger"`
	Recovery RecoveryConfig `yaml:"recovery"`
	Timeout  TimeoutConfig  `yaml:"timeout"`
}

type LoggerConfig struct {
	Enabled bool `yaml:"enabled"`
}

type RecoveryConfig struct {
	Enabled bool `yaml:"enabled"`
}

type TimeoutConfig struct {
	Enabled  bool          `yaml:"enabled"`
	Duration time.Duration `yaml:"duration"`
}

func LoadPath(path string) (string, error) {
	if err := godotenv.Load(path); err != nil {
		return "", err
	}
	cfgPath := os.Getenv(cfgVar)
	if cfgPath == "" {
		return "", errors.New("Failed to get config path")
	}
	return cfgPath, nil
}

func MustLoadByPath(path string) *Config {
	cfgPath, err := LoadPath(path)
	if err != nil {
		panic(err.Error())
	}

	var dbPass string
	dbPass = os.Getenv("DB_PASS")
	if dbPass == "" {
		panic("DB_PASS must be set")
	}

	var cfg Config
	if err = cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		panic("Failed to parse config")
	}

	cfg.DB.Password = dbPass
	return &cfg
}
