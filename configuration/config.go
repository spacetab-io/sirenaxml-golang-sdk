package configuration

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type SirenaConfig struct {
	ClientID                 uint16 `yaml:"client_id,omitempty"`
	Host                     string `yaml:"host,omitempty"`
	Port                     string `yaml:"port,omitempty"`
	KeysPath                 string `yaml:"key_path"`
	ClientPublicKeyFile      string `yaml:"client_public_key_file,omitempty"`
	ClientPrivateKeyFile     string `yaml:"client_private_key_file,omitempty"`
	ClientPrivateKeyPassword string `yaml:"client_private_key_password,omitempty"`
	ServerPublicKeyFile      string `yaml:"server_public_key_file,omitempty"`
	UseSymmetricKeyCrypt     bool   `yaml:"use_symmetric_key_crypt"`
	ZipRequests              bool   `yaml:"zip_requests"`
	ZipResponses             bool   `yaml:"zip_responses"`
	ClientPublicKey          []byte
	ClientPrivateKey         []byte
	ServerPublicKey          []byte
}

// GetSirenaAddr return sirena address to connect client to
func (config *SirenaConfig) GetSirenaAddr() string {
	if config == nil {
		return ""
	}
	if config.Port == "" {
		return config.Host
	}
	return config.Host + ":" + config.Port
}

// GetKeyFile returns contents of key file
func (config *SirenaConfig) GetKeyFile(keyFile string) ([]byte, error) {
	KeyPath := config.KeysPath + "/" + keyFile
	if _, err := os.Stat(KeyPath); os.IsNotExist(err) {
		return nil, err
	}

	return ioutil.ReadFile(KeyPath)
}

func (config *SirenaConfig) GetCerts() (err error) {
	if config.ServerPublicKey, err = config.GetKeyFile(config.ServerPublicKeyFile); err != nil || len(config.ServerPublicKey) == 0 {
		return errors.Wrap(err, "getting server publicKey error")
	}

	if config.ClientPrivateKey, err = config.GetKeyFile(config.ClientPrivateKeyFile); err != nil || len(config.ClientPrivateKey) == 0 {
		return errors.Wrap(err, "getting client privateKey error")
	}

	return
}
