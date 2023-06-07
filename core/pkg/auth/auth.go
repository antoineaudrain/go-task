package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/metadata"
	"os"
	"strings"
	"time"
)

const (
	AccessTokenExpireDuration  = time.Minute * 15
	RefreshTokenExpireDuration = time.Hour * 24
)

const (
	TokenTypeRefresh = "refresh"
	TokenTypeAccess  = "access"
)

type Auth interface {
	GenerateRefreshToken(userId string) (string, error)
	GenerateAccessToken(refreshToken string) (string, error)
	ExtractAccessTokenFromContext(ctx context.Context) string
	Authenticate(ctx context.Context) (string, error)
}

func GenerateRefreshToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userId
	claims["tokenType"] = TokenTypeRefresh
	claims["exp"] = time.Now().Add(RefreshTokenExpireDuration).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateAccessToken(refreshToken string) (string, error) {
	refreshTokenClaims, err := parseJWT(refreshToken)
	if err != nil {
		return "", fmt.Errorf("error parsing refresh token: %v", err)
	}

	userID, ok := refreshTokenClaims["sub"].(string)
	if !ok {
		return "", errors.New("invalid user ID in refresh token")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["tokenType"] = TokenTypeAccess
	claims["exp"] = time.Now().Add(AccessTokenExpireDuration).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %v", err)
	}

	return tokenString, nil
}

func ExtractAccessTokenFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	authorizationHeader := md.Get("authorization")
	if len(authorizationHeader) == 0 {
		return ""
	}

	authValue := strings.TrimSpace(authorizationHeader[0])
	if !strings.HasPrefix(authValue, "Bearer ") {
		return ""
	}

	accessToken := strings.TrimPrefix(authValue, "Bearer ")
	return accessToken
}

func Authenticate(ctx context.Context) (string, error) {
	accessToken := ctx.Value("accessToken").(string)
	if accessToken == "" {
		return "", errors.New("missing access token")
	}

	accessTokenClaims, err := parseJWT(accessToken)
	if err != nil {
		return "", err
	}

	tokenType, ok := accessTokenClaims["tokenType"].(string)
	if !ok || tokenType != TokenTypeAccess {
		return "", errors.New("invalid access token")
	}

	userID, ok := accessTokenClaims["sub"].(string)
	if !ok {
		return "", errors.New("user ID not found in access token claims")
	}

	return userID, nil
}

func parseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid JWT")
	}

	return claims, nil
}
