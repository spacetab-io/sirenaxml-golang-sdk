package sirenaXML

import (
	"math/rand"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	EnvLearning   = "GRU"
	EnvTesting    = "GRT"
	EnvProduction = "GRS"
)

var (
	ipsSlice = []string{"193.104.87.251", "194.84.25.50"}
	portsMap = map[string]string{
		EnvLearning:   "34323",
		EnvTesting:    "34322",
		EnvProduction: "34321",
	}
)

type Config struct {
	ClientID                 uint16 `yaml:"client_id"`
	MaxConnections           uint32 `yaml:"max_connections"`
	Environment              string `yaml:"environment"`
	ClientPublicKey          string `yaml:"client_public_key"`
	ClientPrivateKey         string `yaml:"client_private_key"`
	ClientPrivateKeyPassword string `yaml:"client_private_key_password"`
	ServerPublicKey          string `yaml:"server_public_key"`
	ZippedMessaging          bool   `yaml:"zipped_messaging"`
}

// GetAddr return sirena address to connect client to
func (config *Config) GetAddr() (string, error) {
	if config.Environment == "" {
		return "", errors.New("environment is not set")
	}
	rand.Seed(time.Now().Unix())
	i := rand.Int() % len(ipsSlice)
	return ipsSlice[i] + ":" + portsMap[config.Environment], nil
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

	config.ServerPublicKey = strings.ReplaceAll(config.ServerPublicKey, "\\n", "\n")
	config.ClientPublicKey = strings.ReplaceAll(config.ClientPublicKey, "\\n", "\n")
	config.ClientPrivateKey = strings.ReplaceAll(config.ClientPrivateKey, "\\n", "\n")

	return nil
}
