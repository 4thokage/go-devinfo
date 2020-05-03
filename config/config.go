package config

import (
	"github.com/BurntSushi/toml" // to read config.toml
	"log"
	"os"
)

// Config Represents network port, database server and credentials
type Config struct {
	Port     string
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if len(os.Args) > 1 {
		if _, err := toml.DecodeFile(os.Args[1], &c); err != nil { // try to decode argument file
			log.Fatal(err)
		}
	} else {
		if _, err := toml.DecodeFile("config.toml", &c); err != nil { //decode config file
			log.Fatal(err)
		}
	}
}
