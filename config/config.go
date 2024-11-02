package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerName string            `yaml:"service-name"`
	HTTPServer *HTTPServerConfig `yaml:"http-server"`
	Database   *DatabaseConfig   `yaml:"database"`
	Redis      *RedisConfig      `yaml:"redis"`
}

type HTTPServerConfig struct {
	Address string `yaml:"address"`
}

type DatabaseConfig struct {
	ConnectionString  string `yaml:"connection-string"`
	DbType            string `yaml:"db-type"`
	MigrationFilePath string `yaml:"migration-file-path"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	PoolSize int    `yaml:"pool-size"`
}

func LoadConfig() *Config {

	configPath := os.Getenv("CONFIG_PATH")

	// check configPath in env
	if configPath == "" {
		// get configPath from command
		flagConfigPath := flag.String("config", "", "Path to the configuration file")
		flag.Parse()

		if *flagConfigPath == "" {
			log.Fatal("Configuration file path is required. Use --config=<path>")
		}
		configPath = *flagConfigPath
	}

	var cfg Config

	yamlData, err := os.ReadFile(configPath)

	if err != nil {
		log.Fatal("Error while reading config file ", err)
	}

	yaml.Unmarshal(yamlData, &cfg)

	return &cfg
}
