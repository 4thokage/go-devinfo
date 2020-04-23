package config

import (
	"log" // to print log

	"github.com/BurntSushi/toml" // to read config.toml
)

// Config Represents network port, database server and credentials
type Config struct {
	Port     string
	Server   string
	Database string
}

// Config Reads and parses the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil { // decode file config.toml
		log.Fatal(err)
	}
}
