package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func GenerateBitgetSignature(apiSecret string, apiKey string, passphrase string, method string, uri string, timestamp string) string {

	message := fmt.Sprintf("%s%s%s", timestamp, method, uri)

	// Calculate HMAC-SHA256 signature
	hmac := hmac.New(sha256.New, []byte(apiSecret))
	hmac.Write([]byte(message))
	signature := base64.StdEncoding.EncodeToString(hmac.Sum(nil))

	return signature
}
