package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	// "encoding/hex"
	// "os"
)

// NewRandomKey will generate a 32 bytes random string
func NewRandomKey() []byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		// really, what are you gonna do if randomness failed?
		panic(err)
	}

	return key
}

// Decrypt will return the original value of the encrypted string
func Decrypt(encryptedKey []byte) ([]byte, error) {
	secretKey := getSecret()

	block, err := aes.NewCipher(secretKey)

	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(encryptedKey) < aesgcm.NonceSize() {
		// worth panicking when encrypted key is bad
		panic("Malformed encrypted key")
	}

	return aesgcm.Open(
		nil,
		encryptedKey[:aesgcm.NonceSize()],
		encryptedKey[aesgcm.NonceSize():],
		nil,
	)
}

// Encrypt will encrypt a raw string to
// an encrypted value
// an encrypted value has an IV (nonce) + actual encrypted value
// when we decrypt, we only decrypt the latter part
//encrypty screet key and save to db
func Encrypt(key []byte) ([]byte, error) {
	secretKey := getSecret()

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aesgcm.NonceSize())
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(iv, iv, key, nil)

	return ciphertext, nil
}
//lấy secret key đã lưu vào hệ thống
func getSecret() []byte {
	secret := "chumanhcuong1991"
//os.Getenv("SECRET")
	
	if secret == "" {
		// panic("Error: Must provide a secret key under env variable SECRET")
	}

	secretbite := []byte(secret)

	return secretbite
}

