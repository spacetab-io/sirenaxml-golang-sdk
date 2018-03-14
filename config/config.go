package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/jinzhu/configor"
)

// Config is a config :)
type Config struct {
	LogLevel                 string `yaml:"log_level" env:"LOG_LEVEL" default:"debug"` // log everything by default
	SirenaClientID           string `yaml:"sirena_client_id" env:"SIRENA_CLIENT_ID" required:"true"`
	SirenaHost               string `yaml:"sirena_host" env:"SIRENA_HOST" required:"true"`
	SirenaPort               string `yaml:"sirena_port" env:"SIRENA_PORT" required:"true"`
	ClientPublicKey          string `yaml:"client_public_key" env:"CLIENT_PUBLIC_KEY" required:"true"`
	ClientPrivateKey         string `yaml:"client_private_key" env:"CLIENT_PRIVATE_KEY" required:"true"`
	ClientPrivateKeyPassword string `yaml:"client_private_key_password" env:"CLIENT_PRIVATE_KEY_PASSWORD"`
	ServerPublicKey          string `yaml:"server_public_key" env:"CLIENT_PUBLIC_KEY" required:"true"`
	EnvType                  string
}

var config = &Config{}

// Singleton guard
var once sync.Once

// Get reads config from environment or yaml
func Get() *Config {
	once.Do(func() {
		// Get path to this package
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal("No caller information")
		}
		// Get environment type
		envType := os.Getenv("SIRENA_ENV")
		if envType == "" {
			envType = "dev"
		}
		if err := configor.New(&configor.Config{Environment: envType}).Load(config, path.Dir(filename)+"/config.yaml"); err != nil {
			log.Fatal(err)
		}
		config.EnvType = envType
		configB, err := json.MarshalIndent(config, "", "  ")
		if err == nil {
			fmt.Println("Configuration:", string(configB))
		}
	})
	return config
}

// GetSirenaAddr return sirena address to connect client to
func (config *Config) GetSirenaAddr() string {
	if config == nil {
		return ""
	}
	if config.SirenaPort == "" {
		return config.SirenaHost
	}
	return config.SirenaHost + ":" + config.SirenaPort
}

// GetKeyFile returns contents of key file
func (config *Config) GetKeyFile(keyFile string) ([]byte, error) {
	keyDirs := []string{
		os.Getenv("GOPATH"),
		binaryDir() + "/keys",
	}
	for _, keyDir := range keyDirs {
		exists, err := pathExists(keyDir + "/" + keyFile)
		if err != nil {
			log.Print(err)
		}
		if !exists {
			continue
		}
		return ioutil.ReadFile(keyDir + "/" + keyFile)
	}
	return nil, errors.New("No key files found")
}

// binaryDir returns path where binary was run from
func binaryDir() string {
	ex, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(ex)
}

// pathExists checks if file or dir exist
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
