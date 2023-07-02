package invite_code

import (
	"math/rand"
	"net/url"
	"time"
)

const (
	InvitationCodeLength = 8
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	result := make([]byte, InvitationCodeLength)
	for i := 0; i < InvitationCodeLength; i++ {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func URLEncodeString(input string) string {
	return url.PathEscape(input)
}

func GenerateInvitationCode() string {
	randomString := GenerateRandomString()
	invitationCode := URLEncodeString(randomString)
	return invitationCode
}
