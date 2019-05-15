package token

import (
	"encoding/hex"

	"crypto/hmac"
	"crypto/sha256"
)

// Verify SHA256 hashed signature from Webhook payload
func VerifySignature(messageMAC, message, sharedSecret []byte) bool {
	mac := hmac.New(sha256.New, sharedSecret)
	mac.Write(message)

	expectedMAC := mac.Sum(nil)
	hashedMAC := make([]byte, hex.EncodedLen(len(expectedMAC)))
	hex.Encode(hashedMAC, expectedMAC)

	return hmac.Equal([]byte(messageMAC), hashedMAC)
}