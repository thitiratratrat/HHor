package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
)

type Encryptor interface {
	Encrypt(text string) (string, error)
	Decrypt(text string) (string, error)
}

func EncryptorHandler() Encryptor {
	return &encryptor{
		bytes: []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05},
	}
}

type encryptor struct {
	bytes []byte
}

func (encryptor *encryptor) Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, encryptor.bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (encryptor *encryptor) Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	cipherText, err := base64.StdEncoding.DecodeString(text)

	if err != nil {
		panic(err)
	}

	cfb := cipher.NewCFBDecrypter(block, encryptor.bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}
