package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// Encrypt data using AES
func Encrypt(data string, key []byte) (string, error) {
	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Pad the data to the block size
	padding := aes.BlockSize - len(data)%aes.BlockSize
	data = data + string(make([]byte, padding))

	// Create an initialization vector (IV) for encryption
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Encrypt the data
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(data))

	// Return the encrypted data as a Base64 string
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt the data using AES
func Decrypt(encryptedData string, key []byte) (string, error) {
	// Decode the base64 URL-encoded ciphertext
	ciphertext, err := base64.URLEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	// Create the AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Extract the IV and ciphertext
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Decrypt the data
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	// Return the decrypted data
	return string(ciphertext), nil
}
