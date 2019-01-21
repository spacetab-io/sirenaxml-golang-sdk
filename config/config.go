package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	confpkg "github.com/microparts/configuration-golang"
	yaml "gopkg.in/yaml.v2"
)

// Config is a main Sirena agent config
type Config struct {
	LogLevel                 string `yaml:"log_level,omitempty"`
	SirenaClientID           string `yaml:"sirena_client_id,omitempty"`
	SirenaHost               string `yaml:"sirena_host,omitempty"`
	SirenaPort               string `yaml:"sirena_port,omitempty"`
	ClientPublicKey          string `yaml:"client_public_key,omitempty"`
	ClientPrivateKey         string `yaml:"client_private_key,omitempty"`
	ClientPrivateKeyPassword string `yaml:"client_private_key_password,omitempty"`
	ServerPublicKey          string `yaml:"server_public_key,omitempty"`
}

// CNFG is a Config singletone
var CNFG *Config

func init() {
	CNFG = loadConfig()
}

// Get returns config
func Get() *Config {
	return CNFG
}

// loadConfig loads config from YAML files
func loadConfig() *Config {
	configPath := confpkg.GetEnv("CONFIG_PATH", "config")
	fmt.Printf("Config path: %s\n", configPath)

	configBytes, err := confpkg.ReadConfigs(configPath)
	if err != nil {
		log.Fatalf("Error reading config: %+v", err)
	}
	fmt.Printf("Config contents:\n%s\n", configBytes)

	var config Config
	yaml.Unmarshal(configBytes, &config)

	return &config
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

// KeyDirs is a list of directories to search Sirena key files in
var KeyDirs = []string{
	os.Getenv("GOPATH"),
	binaryDir() + "/keys",
	pwdDir() + "/sirena-agent-go/keys",
	pwdDir() + "/sirena-proxy/keys",
}

// GetKeyFile returns contents of key file
func (config *Config) GetKeyFile(keyFile string) ([]byte, error) {
	for _, keyDir := range KeyDirs {
		exists, err := pathExists(keyDir + "/" + keyFile)
		if err != nil {
			log.Print(err)
		}
		if !exists {
			continue
		}
		return ioutil.ReadFile(keyDir + "/" + keyFile)
	}
	return nil, fmt.Errorf("No key file %s found", keyFile)
}

// binaryDir returns path where binary was run from
func binaryDir() string {
	ex, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(ex)
}

// pwdDir returns pwd dir
func pwdDir() string {
	ex, err := os.Getwd()
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
