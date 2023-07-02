package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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
	ExtractAuthenticationToken(authHeader string) string
	Authenticate(ctx context.Context) (string, error)
	ParseJWT(tokenString string) (jwt.MapClaims, error)
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
	refreshTokenClaims, err := ParseJWT(refreshToken)
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

func ExtractAuthenticationToken(authHeader string) string {
	accessToken := ""
	if strings.HasPrefix(authHeader, "Bearer ") {
		accessToken = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		accessToken = authHeader
	}
	return accessToken
}

func ValidateAccessToken(token string) (*uuid.UUID, error) {
	accessTokenClaims, err := ParseJWT(token)

	tokenType, ok := accessTokenClaims["tokenType"].(string)
	if !ok || tokenType != TokenTypeAccess {
		return nil, errors.New("invalid access token type")
	}

	userID, ok := accessTokenClaims["sub"].(string)
	if !ok || userID == "" {
		return nil, errors.New("user ID not found in access token")
	}

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid access token user ID")
	}

	return &parsedUserID, nil
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
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

func Authenticate(ctx context.Context) (uuid.UUID, error) {
	accessToken := ctx.Value("accessToken").(string)
	if accessToken == "" {
		return uuid.UUID{}, errors.New("missing access token")
	}

	accessTokenClaims, err := parseJWT(accessToken)
	if err != nil {
		return uuid.UUID{}, err
	}

	tokenType, ok := accessTokenClaims["tokenType"].(string)
	if !ok || tokenType != TokenTypeAccess {
		return uuid.UUID{}, errors.New("invalid access token")
	}

	userID, ok := accessTokenClaims["sub"].(string)
	if !ok || userID == "" {
		return uuid.UUID{}, errors.New("user ID not found in access token claims")
	}

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return uuid.UUID{}, errors.New("invalid user ID")
	}

	return parsedUserID, nil
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
