package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config : Represents database server and credentials
type Config struct {
	Server   string
	Database string
	Port     string
}

// Range : Represents the color range
type Range struct {
	Blue   []int
	Green  []int
	Orange []int
	Red    []int
}

// Read and parse the configuration file for credentials
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

// Read and parse the configuration file for range
func (r *Range) Read() {
	if _, err := toml.DecodeFile("config.toml", &r); err != nil {
		log.Fatal(err)
	}
}
