package sirena

import (
	"strings"

	"github.com/pkg/errors"
)

type Config struct {
	ClientID                 uint16 `yaml:"client_id"`
	RequestHandlers          uint32 `yaml:"putls_count"`
	Ip                       string `yaml:"ip"`
	Port                     string `yaml:"port"`
	KeysPath                 string `yaml:"key_path"`
	ClientPublicKey          []byte `yaml:"client_public_key"`
	ClientPrivateKey         []byte `yaml:"client_private_key"`
	ClientPrivateKeyPassword string `yaml:"client_private_key_password"`
	ServerPublicKey          []byte `yaml:"server_public_key"`
	ZippedMessaging          bool   `yaml:"zipped_messaging"`
}

// GetAddr return sirena address to connect client to
func (config *Config) GetAddr() string {
	if config == nil {
		return ""
	}
	if config.Port == "" {
		return config.Ip
	}
	return config.Ip + ":" + config.Port
}

func (config *Config) PrepareKeys() error {
	if len(config.ServerPublicKey) == 0 {
		return errors.New("server public key is empty")
	}
	if len(config.ClientPublicKey) == 0 {
		return errors.New("client public key is empty")
	}
	if len(config.ClientPrivateKey) == 0 {
		return errors.New("client private key is empty")
	}

	config.ServerPublicKey = []byte(strings.ReplaceAll(string(config.ServerPublicKey), "\\n", "\n"))
	config.ClientPublicKey = []byte(strings.ReplaceAll(string(config.ClientPublicKey), "\\n", "\n"))
	config.ClientPrivateKey = []byte(strings.ReplaceAll(string(config.ClientPrivateKey), "\\n", "\n"))

	return nil
}
