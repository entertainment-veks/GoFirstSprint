package config

import (
	"flag"
	"os"
	"sync"
)

const (
	serverAddressKey   = "SERVER_ADDRESS"
	baseURLKey         = "BASE_URL_KEY"
	fileStoragePathKey = "FILE_STORAGE_PATH"
	databaseDSNKey     = "DATABASE_DSN"
)

const (
	defServerAddress   = ":8080"
	defBaseURL         = "http://127.0.0.1:8080"
	defFileStoragePath = "file"
	defDatabaseURL     = "postgres://postgres:postgres@localhost:5432/shortener?sslmode=disable"
)

var (
	configuredByFlags sync.Once
)

type Config struct {
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
	DatabaseURL     string
}

var (
	ServerAddressFlagsValue   string
	BaseURLFlagsValue         string
	FileStoragePathFlagsValue string
	DatabaseURLFlagsValue     string
)

func NewConfig() *Config {
	c := &Config{
		ServerAddress:   defServerAddress,
		BaseURL:         defBaseURL,
		FileStoragePath: defFileStoragePath,
		DatabaseURL:     defDatabaseURL,
	}

	c.configureViaEnv()
	c.configureViaFlags()

	return c
}

func (c *Config) configureViaEnv() {
	if val := os.Getenv(serverAddressKey); len(val) != 0 {
		c.ServerAddress = val
	}

	if val := os.Getenv(baseURLKey); len(val) != 0 {
		c.BaseURL = val
	}

	if val := os.Getenv(fileStoragePathKey); len(val) != 0 {
		c.FileStoragePath = val
	}

	if val := os.Getenv(databaseDSNKey); len(val) != 0 {
		c.DatabaseURL = val
	}
}

func (c *Config) configureViaFlags() {
	configuredByFlags.Do(func() {
		flag.Func("a", "Server address", func(s string) error {
			ServerAddressFlagsValue = s
			return nil
		})

		flag.Func("b", "Base url", func(s string) error {
			BaseURLFlagsValue = s
			return nil
		})

		flag.Func("f", "File storage path", func(s string) error {
			FileStoragePathFlagsValue = s
			return nil
		})

		flag.Func("d", "Database path", func(s string) error {
			DatabaseURLFlagsValue = s
			return nil
		})

		flag.Parse()
	})

	if len(ServerAddressFlagsValue) != 0 {
		c.ServerAddress = ServerAddressFlagsValue
	}
	if len(BaseURLFlagsValue) != 0 {
		c.BaseURL = BaseURLFlagsValue
	}
	if len(FileStoragePathFlagsValue) != 0 {
		c.FileStoragePath = FileStoragePathFlagsValue
	}
	if len(DatabaseURLFlagsValue) != 0 {
		c.DatabaseURL = DatabaseURLFlagsValue
	}
}
