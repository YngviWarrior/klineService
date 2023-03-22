package utils

import (
	rand "crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
)

func GerHash(length int64) string {
	hasher := sha1.New()

	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)

	if err != nil {
		panic(err)
	}

	random := base64.StdEncoding.EncodeToString(randomBytes)

	hasher.Write([]byte(random))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))[:length]
}

func EncryptPassHS256(pass string) (sha string) {
	hasher := sha256.New()
	hasher.Write([]byte(pass))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
