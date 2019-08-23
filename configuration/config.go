package sirenaXML

import (
	"net"
	"strings"

	"github.com/pkg/errors"
)

const (
	EnvLearning   = "GRU"
	EnvTesting    = "GRT"
	EnvProduction = "GRS"
)

var (
	portsMap = map[string]int{
		EnvLearning:   34323,
		EnvTesting:    34322,
		EnvProduction: 34321,
	}
)

type Config struct {
	ClientID                 uint16 `yaml:"client_id"`
	MaxConnections           uint32 `yaml:"max_connections"`
	Ip                       string `yaml:"ip"`
	Environment              string `yaml:"environment"`
	ClientPublicKey          string `yaml:"client_public_key"`
	ClientPrivateKey         string `yaml:"client_private_key"`
	ClientPrivateKeyPassword string `yaml:"client_private_key_password"`
	ServerPublicKey          string `yaml:"server_public_key"`
	ZippedMessaging          bool   `yaml:"zipped_messaging"`
	MaxConnectTries          int    `yaml:"max_connect_tries"`
}

// GetAddr return sirena address to connect client to
func (c *Config) GetAddr() (*net.TCPAddr, error) {
	if c.Environment == "" {
		return nil, errors.New("environment is not set")
	}
	if c.Ip == "" {
		return nil, errors.New("ip is not set")
	}
	return &net.TCPAddr{Port: portsMap[c.Environment], IP: net.ParseIP(c.Ip)}, nil
}

func (c *Config) PrepareKeys() error {
	if len(c.ServerPublicKey) == 0 {
		return errors.New("server public key is empty")
	}
	if len(c.ClientPublicKey) == 0 {
		return errors.New("client public key is empty")
	}
	if len(c.ClientPrivateKey) == 0 {
		return errors.New("client private key is empty")
	}

	c.ServerPublicKey = strings.ReplaceAll(c.ServerPublicKey, "\\n", "\n")
	c.ClientPublicKey = strings.ReplaceAll(c.ClientPublicKey, "\\n", "\n")
	c.ClientPrivateKey = strings.ReplaceAll(c.ClientPrivateKey, "\\n", "\n")

	return nil
}

type KeyType int

const (
	ServerPublicKey KeyType = iota
	ClientPublicKey
	ClientPrivateKey
)

type KeysData struct {
	Keys              map[KeyType][]byte
	ClientPrivKeyPass string
}

func (c *Config) GetKeys() *KeysData {
	return &KeysData{
		Keys: map[KeyType][]byte{
			ServerPublicKey:  []byte(c.ServerPublicKey),
			ClientPrivateKey: []byte(c.ClientPrivateKey),
			ClientPublicKey:  []byte(c.ClientPublicKey),
		},
		ClientPrivKeyPass: c.ClientPrivateKeyPassword,
	}
}
