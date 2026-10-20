package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type PayLoad struct {
	Sub         string `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func Base64ByteArrayEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

func CreateJwt(secret string, data PayLoad) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	// 1. Convert header and payload to JSON bytes
	byteArrayHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	byteArrayPayLoad, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// 2. Base64URL-encode header and payload
	headerBase64 := Base64ByteArrayEncode(byteArrayHeader)
	payLoadBase64 := Base64ByteArrayEncode(byteArrayPayLoad)

	// 3. Create the unsigned message: header.payload
	unsignedMessage := headerBase64 + "." + payLoadBase64

	// 4. Compute HMAC-SHA256 signature
	hmacSha256 := hmac.New(sha256.New, []byte(secret))
	hmacSha256.Write([]byte(unsignedMessage))
	signature := hmacSha256.Sum(nil)

	signatureBase64 := Base64ByteArrayEncode(signature)

	// 5. Assemble final JWT
	jwt := unsignedMessage + "." + signatureBase64 // signedMessage(jwt) = headerBase64 + "." + payLoadBase64 + "." + signatureBase64

	return jwt, nil
}
