package aesenc

import "github.com/gogf/gf/crypto/gaes"

type Encryption struct {
	key []byte
}

func New(key string) *Encryption {
	return &Encryption{key: []byte(key)}
}

func (Self *Encryption) Encrypt(data []byte) ([]byte, error) {
	return gaes.Encrypt(data, Self.key)
}

func (Self *Encryption) Decrypt(data []byte) ([]byte, error) {
	return gaes.Decrypt(data, Self.key)
}
