package config

import (
	"egghead/app/util"
)

type Database struct {
	DatabaseHost     string
	DatabasePort     int
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseSSLMode  bool
}

type Service struct {
	ServerPort  string
	Environment string
	ServiceName string
}

// AppConfig represents the configuration for your application
type Config struct {
	DatabaseHost      string
	DatabasePort      int
	DatabaseName      string
	DatabaseUser      string
	DatabasePassword  string
	DatabaseSSLMode   string
	DatabaseServerCA  string
	DatabaseClientCA  string
	DatabaseClientKey string
	ServerPort        string
	Environment       EnvironmentEnum
	ServiceName       string
	Debug             bool
}

// SSL Mode for the postgres connection
type PGSSLMode string

const (
	Require    PGSSLMode = "require"
	Disable    PGSSLMode = "disable"
	Allow      PGSSLMode = "allow"
	Prefer     PGSSLMode = "prefer"
	VerifyCA   PGSSLMode = "verify-ca"
	VerifyFull PGSSLMode = "verify-full"
)

// Environment is a custom type for the environment enum
type EnvironmentEnum string

// TransactionType enum
const (
	Local      EnvironmentEnum = "local"
	Develop    EnvironmentEnum = "development"
	Production EnvironmentEnum = "production"
	Staging    EnvironmentEnum = "staging"
)

// NewAppConfig creates a new AppConfig instance with values initialized from environment variables
func LoadConfig() *Config {
	config := &Config{}

	config.DatabaseHost = util.GetEnv("DATABASE_HOST", "localhost")
	config.DatabasePort = util.ConvertStrToInt(util.GetEnv("DATABASE_PORT", "5432"))
	config.DatabaseName = util.GetEnv("DATABASE_NAME", "egghead")
	config.DatabaseUser = util.GetEnv("DATABASE_USER", "admin")
	config.DatabasePassword = util.GetEnv("DATABASE_PASSWORD", "admin")
	config.DatabaseSSLMode = util.GetEnv("DATABASE_SSL_MODE", "disable")
	config.DatabaseServerCA = util.GetEnv("DATABASE_SERVER_CA", "")
	config.DatabaseClientCA = util.GetEnv("DATABASE_CLIENT_CA", "")
	config.DatabaseClientKey = util.GetEnv("DATABASE_CLIENT_KEY", "")
	config.ServerPort = util.GetEnv("SERVER_PORT", "8000")
	config.ServiceName = util.GetEnv("SERVICE_NAME", "egghead")
	config.Environment = EnvironmentEnum(util.GetEnvWithEnum("ENV", string(Local), []string{string(Local), string(Develop), string(Production)}))
	config.Debug = util.GetEnvAsBool("DEBUG", false)

	return config
}
