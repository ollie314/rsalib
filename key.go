package rsalib

import (
  "errors"
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509"
  "encoding/pem"
)

const (
  privateKeyType string = "RSA PRIVATE KEY"
  publicKeyType string = "RSA PUBLIC KEY"

  pemFormat string = "PEM"
)

type KeyPair struct {
  Format string
  PrivateKey string
  PublicKey string
}

// Generate a random private key with given length
func generatePrivateKey(lenght int) (*rsa.PrivateKey, error) {
  key, err := rsa.GenerateKey(rand.Reader, lenght); if err != nil {
    // TODO: log
    return nil, err
  }
  return key, nil;
}

func privateKeyToPEM(key *rsa.PrivateKey, keyType string, pwd string) (*pem.Block, error) {
  block := &pem.Block {
      Type:  keyType,
      Bytes: x509.MarshalPKCS1PrivateKey(key),
  }
  if pwd != "" {
      b, err := x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(pwd), x509.PEMCipherAES256)
      if err != nil {
          return b, err
      }
  }
  return block, nil
}

func encodeToString(block *pem.Block) (string) {
  return string(pem.EncodeToMemory(block))
}

func publicKeyToPEM(key *rsa.PrivateKey, keyType string) (*pem.Block, error) {
  return &pem.Block {
    Type: keyType,
    Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey),
  }, nil
}

func GenerateKeyPair(length int, pwd string) (*KeyPair, error) {
  priv, err := generatePrivateKey(length); if err != nil {
    return nil, errors.New("error occurred during private key generation process")
  }
  privPem, err := privateKeyToPEM(priv, privateKeyType, pwd); if err != nil {
    return nil, errors.New("error occurred during private key pem generation")
  }
  pubPem, err := publicKeyToPEM(priv, publicKeyType); if err != nil {
    return nil, errors.New("error occurred during public key pem generation")
  }
  return &KeyPair{pemFormat, encodeToString(privPem), encodeToString(pubPem)}, nil
}
