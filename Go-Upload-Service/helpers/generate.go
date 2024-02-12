package helpers

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString() (string, error) {
	
    byteSize := (4 * 7 / 3) + 2
    bytes := make([]byte, byteSize)


    _, err := rand.Read(bytes)
    if err != nil {
        return "", err
    }

    randomString := base64.URLEncoding.EncodeToString(bytes)


    randomString = randomString[:7]

    return randomString, nil
}
