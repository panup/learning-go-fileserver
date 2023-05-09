package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
)

func encryptBytes(data []byte, encryptionKey []byte) ([]byte, error) {

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	enrypted := gcm.Seal(nonce, nonce, data, nil)

	return enrypted, nil

}

func decryptBytes(encrypteddata []byte, encryptionKey []byte) ([]byte, error) {

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	nonceSize := gcm.NonceSize()
	nonce, ciphered := encrypteddata[:nonceSize], encrypteddata[nonceSize:]

	derypted, err := gcm.Open(nil, nonce, ciphered, nil)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	return derypted, nil

}

func generateKey() string {
	key := strings.Replace(uuid.New().String(), "-", "", -1)

	return key
}
