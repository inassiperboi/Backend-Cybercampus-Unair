package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"cybercampus_module/configs"
	"cybercampus_module/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func GenerateToken(id string, username string, email string, jenisUser string, role string) (string, error) {

	header := models.Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	expirationTime := time.Now().Add(1 * time.Hour).Unix()

	payload := models.PayLoad{
		ID:       id,
		Username: username,
		Email:    email,
		JenisUser: jenisUser,
		Role:     role,
		Exp:      expirationTime,
	}

	headerBytes, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	headerEncoded := base64.RawURLEncoding.EncodeToString(headerBytes)
	payloadEncoded := base64.RawURLEncoding.EncodeToString(payloadBytes)

	secretKey := os.Getenv(configs.LoadEnv("SECRET_KEY"))
	signature := CreateSignature(headerEncoded, payloadEncoded, secretKey)

	token := fmt.Sprintf("%s.%s.%s", headerEncoded, payloadEncoded, signature)

	return token, nil

}

func CreateSignature(header string, payload string, secret string) string {
	data := hmac.New(sha256.New, []byte(secret))
	data.Write([]byte(fmt.Sprintf("%s.%s", header, payload)))
	signature := base64.RawURLEncoding.EncodeToString(data.Sum(nil))
	return signature
}