package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func Hash(data []byte) (string, error) {
	hasher := sha1.New()
	_, err := hasher.Write(data)
	if err != nil {
		return "", fmt.Errorf("hash error: %w", err)
	}

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil)), nil
}
