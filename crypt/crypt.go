package crypt

import (
	"crypto"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func EncryptDataWithServerPubKey(data, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("Failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("Failed to cast public key into *rsa.PublicKey")
	}
	encryptedData, err := rsa.EncryptPKCS1v15(crand.Reader, pubKey, data)
	if err != nil {
		return nil, err
	}
	return encryptedData, nil
}

func DecryptDataWithClientPrivateKey(data, key []byte, keyPassword string) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("Failed to decode PEM block containing private key")
	}
	var privateKey *rsa.PrivateKey
	if keyPassword != "" {
		key, err := x509.DecryptPEMBlock(block, []byte(keyPassword))
		if err != nil {
			return nil, err
		}
		privateKey, err = x509.ParsePKCS1PrivateKey(key)
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	}
	encryptedData, err := rsa.DecryptPKCS1v15(crand.Reader, privateKey, data)
	if err != nil {
		return nil, err
	}
	return encryptedData, nil
}

func GeneratePrivateKeySignature(data, key []byte, keyPassword string) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("Failed to decode PEM block containing private key")
	}
	var privateKey *rsa.PrivateKey
	if keyPassword != "" {
		key, err := x509.DecryptPEMBlock(block, []byte(keyPassword))
		if err != nil {
			return nil, err
		}
		privateKey, err = x509.ParsePKCS1PrivateKey(key)
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	}
	h := sha1.New()
	h.Write(data)
	digest := h.Sum(nil)
	s, err := rsa.SignPKCS1v15(crand.Reader, privateKey, crypto.SHA1, digest)
	if err != nil {
		return nil, err
	}
	return s, nil
}
