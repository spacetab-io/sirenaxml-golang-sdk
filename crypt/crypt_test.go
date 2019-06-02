package crypt

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	privateKey,
	publicKey,
	privateKeySecret,
	publicKeySecret []byte
	pass = "secret"
	key  = []byte("deskey")
)

func tearUp() {
	var err error
	privateKey, err = ioutil.ReadFile("private.key.pem")
	if err != nil {
		panic(err)
	}
	publicKey, err = ioutil.ReadFile("public.key.pem")
	if err != nil {
		panic(err)
	}

	privateKeySecret, err = ioutil.ReadFile("private.secret.key.pem")
	if err != nil {
		panic(err)
	}
	publicKeySecret, err = ioutil.ReadFile("public.secret.key.pem")
	if err != nil {
		panic(err)
	}
	publicKeySecret, err = ioutil.ReadFile("public.secret.key.pem")
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	tearUp()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestEncryptDataWithServerPubKey(t *testing.T) {
	data := []byte(RandString(8))
	t.Run("success", func(t *testing.T) {
		encKey, err := EncryptDataWithServerPubKey(data, publicKey)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.NotEmpty(t, encKey)
		assert.NotEqual(t, data, encKey)
	})
	t.Run("error on empty key", func(t *testing.T) {
		_, err := EncryptDataWithServerPubKey([]byte("some"), nil)
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
	t.Run("error on unknown key type (not RSA, DSA or ECDSA)", func(t *testing.T) {
		key := []byte(strings.ReplaceAll("1-----BEGIN PUBLIC KEY-----\nbleah\n-----END PUBLIC KEY-----", "\\n", "\n"))
		_, err := EncryptDataWithServerPubKey([]byte("some"), key)
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
}

func TestDecryptDataWithClientPrivateKey(t *testing.T) {
	data := []byte(RandString(8))
	//serverPubKey := []byte(strings.ReplaceAll(os.Getenv("CLIENT_PUBLIC_KEY"), "\\n", "\n"))
	//clientPrivateKey := []byte(strings.ReplaceAll(os.Getenv("CLIENT_PRIVATE_KEY"), "\\n", "\n"))
	//pass := os.Getenv("CLIENT_PRIVATE_KEY_PASSWORD")
	t.Run("success no password", func(t *testing.T) {
		encData, err := EncryptDataWithServerPubKey(data, publicKey)
		if err != nil {
			t.FailNow()
		}
		decData, err := DecryptDataWithClientPrivateKey(encData, privateKey, "")
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.Equal(t, data, decData)
	})
	t.Run("success with password", func(t *testing.T) {
		encData, err := EncryptDataWithServerPubKey(data, publicKeySecret)
		if err != nil {
			t.FailNow()
		}
		decData, err := DecryptDataWithClientPrivateKey(encData, privateKeySecret, pass)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.Equal(t, data, decData)
	})
	t.Run("error with password", func(t *testing.T) {
		encData, _ := EncryptDataWithServerPubKey(data, publicKey)
		_, err := DecryptDataWithClientPrivateKey(encData, privateKey, "wrongSecret")
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
}

func TestGeneratePrivateKeySignature(t *testing.T) {
	deskeyCrypted, err := EncryptDataWithServerPubKey(key, publicKey)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	_, err = GeneratePrivateKeySignature(deskeyCrypted, privateKey, "")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
}
