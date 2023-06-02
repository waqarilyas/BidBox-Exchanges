package bybit

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GenerateBybitSignature(apiKey, apiSecret string, recvWindow, timestamp int64, queryString string) string {
	dataToSign := fmt.Sprintf("%d%s%d%s", timestamp, apiKey, recvWindow, queryString)
	hmacKey := []byte(apiSecret)
	hmacHash := hmac.New(sha256.New, hmacKey)
	hmacHash.Write([]byte(dataToSign))
	signature := hex.EncodeToString(hmacHash.Sum(nil))
	return signature
}

func GenerateBybitV2Signature(queryString, apiSecret string) string {
	hmacHash := hmac.New(sha256.New, []byte(apiSecret))
	hmacHash.Write([]byte(queryString))
	signature := hmacHash.Sum(nil)

	return hex.EncodeToString(signature)
}
