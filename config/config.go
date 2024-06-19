package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	HttpServer HttpServer     `yaml:"server"`
	Database   DataBaseConfig `yaml:"database"`
	LogLevel   string
}

func (c Config) String() string {
	return fmt.Sprintf("host: %s ; port: %d; dbUser: %s ; password: %s ; name: %s",
		c.Database.Host, c.Database.Port, c.Database.DbUser, c.Database.Password, c.Database.Name)
}

type DataBaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbUser   string `yaml:"dbUser"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type HttpServer struct {
	Port int `yaml:"port"`
}

func (c *Config) Validate() error {
	if c.HttpServer.Port <= 0 || c.HttpServer.Port > 65353 {
		return fmt.Errorf("invalid port for http server: %d", c.HttpServer.Port)
	}
	if c.Database.Port <= 0 || c.HttpServer.Port > 65353 {
		return fmt.Errorf("invalid port for http server: %d", c.Database.Port)
	}
	if c.Database.Host == "" {
		return fmt.Errorf("database host is empty!")
	}
	if c.Database.DbUser == "" {
		return fmt.Errorf("database user is empty!")
	}
	if c.Database.Name == "" {
		return fmt.Errorf("database name is empty!")
	}

	return nil
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var config Config
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding YAML file: %v", err)
	}

	res := config.String()
	fmt.Println(res)

	return &config, nil
}
