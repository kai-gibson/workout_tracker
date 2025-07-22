package config

import (
	"os"
	"strconv"
)

/* Keep values in config struct tightly sized for lower memory usage */
type config struct {
	dbDsn       string
	skipMigrate bool
	devMode     bool
	port        int32
	logLevel    string
}

// Read only interface for config
// You can copy paste the config struct fields then run this in vim to generate:
// s/\v\s*(\S)(\S+)\s*(\S+)/\t\U\1\E\2\(\) \3/g
// You could make a regex replace to make the func defs too...

type Config interface {
	DbDsn() string
	SkipMigrate() bool
	DevMode() bool
	Port() int32
	LogLevel() string
}

func (cfg *config) DbDsn() string {
	return cfg.dbDsn
}

func (cfg *config) SkipMigrate() bool {
	return cfg.skipMigrate
}

func (cfg *config) DevMode() bool {
	return cfg.devMode
}

func (cfg *config) Port() int32 {
	return cfg.port
}

func (cfg *config) LogLevel() string {
	return cfg.logLevel
}

func NewFromEnv() (Config, error) {
	config := config{
		dbDsn: os.Getenv("DBDSN"),
    logLevel: os.Getenv("LOGLEVEL"),
	}

	var err error

	config.skipMigrate, err = strconv.ParseBool(os.Getenv("SKIPMIGRATE"))
	if err != nil {
		return nil, err
	}

	config.devMode, err = strconv.ParseBool(os.Getenv("DEVMODE"))
	if err != nil {
		return nil, err
	}

	intVal, err := strconv.ParseUint(os.Getenv("PORT"), 10, 32)
	if err != nil {
		return nil, err
	}
	config.port = int32(intVal)

	return &config, nil
}
