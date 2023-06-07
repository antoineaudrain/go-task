package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	customErrors "go-task/core/pkg/errors"
	mocksCore "go-task/core/pkg/mocks"
	"go-task/core/pkg/models"
	"go-task/user/internal/user/mocks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "go-task/user/api"
)

func TestHandler_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockService := new(mocks.Service)
		handler := &Handler{
			service: mockService,
		}

		req := &pb.CreateUserRequest{
			Email:    "test@example.com",
			Password: "testPassword",
			FullName: "Test User",
		}

		expectedUser := &models.User{
			ID:           uuid.New(),
			Email:        req.GetEmail(),
			PasswordHash: "hashedPassword",
			FullName:     req.GetFullName(),
		}

		mockService.EXPECT().CreateUserWithHashedPassword(req.GetEmail(), req.GetPassword(), req.GetFullName()).Return(expectedUser, nil)

		resp, err := handler.CreateUser(context.Background(), req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, expectedUser.ID.String(), resp.GetUser().GetId())
		assert.Equal(t, expectedUser.Email, resp.GetUser().GetEmail())
		assert.Equal(t, expectedUser.FullName, resp.GetUser().GetFullName())
	})

	t.Run("hashing_error", func(t *testing.T) {
		mockService := new(mocks.Service)
		handler := &Handler{
			service: mockService,
		}

		req := &pb.CreateUserRequest{
			Email:    "test@example.com",
			Password: "testPassword",
			FullName: "Test User",
		}

		mockService.EXPECT().CreateUserWithHashedPassword(req.GetEmail(), req.GetPassword(), req.GetFullName()).Return(nil, customErrors.NewHashingError("failed to hash password", errors.New("hashing error")))

		resp, err := handler.CreateUser(context.Background(), req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("database_error_user_already_exists", func(t *testing.T) {
		mockService := new(mocks.Service)
		handler := &Handler{
			service: mockService,
		}

		req := &pb.CreateUserRequest{
			Email:    "test@example.com",
			Password: "testPassword",
			FullName: "Test User",
		}

		mockError := &pgconn.PgError{Code: "23505"}
		mockService.EXPECT().CreateUserWithHashedPassword(req.GetEmail(), req.GetPassword(), req.GetFullName()).Return(nil, customErrors.NewDatabaseError("failed to create user", mockError))

		resp, err := handler.CreateUser(context.Background(), req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("database_error_other", func(t *testing.T) {
		mockService := new(mocks.Service)
		handler := &Handler{
			service: mockService,
		}

		req := &pb.CreateUserRequest{
			Email:    "test@example.com",
			Password: "testPassword",
			FullName: "Test User",
		}

		mockError := errors.New("database error")
		mockService.EXPECT().CreateUserWithHashedPassword(req.GetEmail(), req.GetPassword(), req.GetFullName()).Return(nil, customErrors.NewDatabaseError("failed to create user", mockError))

		resp, err := handler.CreateUser(context.Background(), req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Internal, status.Code(err))
	})
}

func TestHandler_Login(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockService := new(mocks.Service)
		mockAuth := new(mocksCore.Auth)
		handler := &Handler{
			service: mockService,
		}

		req := &pb.LoginRequest{
			Email:    "test@example.com",
			Password: "testPassword",
		}

		expectedUser := &models.User{
			ID:       uuid.New(),
			Email:    req.GetEmail(),
			FullName: "Test User",
		}

		expectedRefreshToken := "refreshToken"
		expectedAccessToken := "accessToken"

		mockService.EXPECT().AuthenticateUser(req.GetEmail(), req.GetPassword()).Return(expectedUser, nil)
		mockAuth.EXPECT().GenerateRefreshToken(expectedUser.ID.String()).Return(expectedRefreshToken, nil)
		mockAuth.EXPECT().GenerateAccessToken(expectedRefreshToken).Return(expectedAccessToken, nil)

		resp, err := handler.Login(context.Background(), req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, expectedUser.ID.String(), resp.GetUser().GetId())
		assert.Equal(t, expectedUser.Email, resp.GetUser().GetEmail())
		assert.Equal(t, expectedUser.FullName, resp.GetUser().GetFullName())
		assert.Equal(t, expectedRefreshToken, resp.GetRefreshToken())
		assert.Equal(t, expectedAccessToken, resp.GetAccessToken())
	})

	t.Run("invalid credentials", func(t *testing.T) {
		mockService := new(mocks.Service)
		handler := &Handler{
			service: mockService,
		}

		req := &pb.LoginRequest{
			Email:    "test@example.com",
			Password: "testPassword",
		}

		mockService.EXPECT().AuthenticateUser(req.GetEmail(), req.GetPassword()).Return(nil, errors.New("invalid credentials"))

		resp, err := handler.Login(context.Background(), req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Unauthenticated, status.Code(err))
	})

	t.Run("failed authentication", func(t *testing.T) {
		mockService := new(mocks.Service)
		handler := &Handler{
			service: mockService,
		}

		req := &pb.LoginRequest{
			Email:    "test@example.com",
			Password: "testPassword",
		}

		mockService.EXPECT().AuthenticateUser(req.GetEmail(), req.GetPassword()).Return(nil, errors.New("authentication failed"))

		resp, err := handler.Login(context.Background(), req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Unauthenticated, status.Code(err))
	})

	t.Run("failed refresh token generation", func(t *testing.T) {
		mockService := new(mocks.Service)
		mockAuth := new(mocksCore.Auth)
		handler := &Handler{
			service: mockService,
		}

		req := &pb.LoginRequest{
			Email:    "test@example.com",
			Password: "testPassword",
		}

		expectedUser := &models.User{
			ID:       uuid.New(),
			Email:    req.GetEmail(),
			FullName: "Test User",
		}

		mockService.EXPECT().AuthenticateUser(req.GetEmail(), req.GetPassword()).Return(expectedUser, nil)
		mockAuth.EXPECT().GenerateRefreshToken(expectedUser.ID.String()).Return("", errors.New("failed to generate refresh token"))

		resp, err := handler.Login(context.Background(), req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("failed access token generation", func(t *testing.T) {
		mockService := new(mocks.Service)
		mockAuth := new(mocksCore.Auth)
		handler := &Handler{
			service: mockService,
		}

		req := &pb.LoginRequest{
			Email:    "test@example.com",
			Password: "testPassword",
		}

		expectedUser := &models.User{
			ID:       uuid.New(),
			Email:    req.GetEmail(),
			FullName: "Test User",
		}

		expectedRefreshToken := "refreshToken"

		mockService.EXPECT().AuthenticateUser(req.GetEmail(), req.GetPassword()).Return(expectedUser, nil)
		mockAuth.EXPECT().GenerateRefreshToken(expectedUser.ID.String()).Return(expectedRefreshToken, nil)
		mockAuth.EXPECT().GenerateAccessToken(expectedRefreshToken).Return("", errors.New("failed to generate access token"))

		resp, err := handler.Login(context.Background(), req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Internal, status.Code(err))
	})
}

func TestHandler_RefreshToken(t *testing.T) {
	//t.Run("success", func(t *testing.T) {
	//	mockAuth := new(mocks.Auth)
	//	handler := &Handler{
	//		auth: mockAuth,
	//	}
	//
	//	req := &pb.RefreshTokenRequest{
	//		RefreshToken: "validRefreshToken",
	//	}
	//
	//	expectedAccessToken := "validAccessToken"
	//
	//	mockAuth.EXPECT().GenerateAccessToken(req.GetRefreshToken()).Return(expectedAccessToken, nil)
	//
	//	resp, err := handler.RefreshToken(context.Background(), req)
	//
	//	assert.NoError(t, err)
	//	assert.NotNil(t, resp)
	//	assert.Equal(t, expectedAccessToken, resp.GetAccessToken())
	//})

	//t.Run("invalid refreshToken", func(t *testing.T) {
	//	mockAuth := new(mocks.Auth)
	//	handler := &Handler{
	//		auth: mockAuth,
	//	}
	//
	//	resp, err := handler.RefreshToken(context.Background(), nil)
	//
	//	assert.Error(t, err)
	//	assert.Nil(t, resp)
	//	assert.Equal(t, codes.InvalidArgument, status.Code(err))
	//	assert.Contains(t, err.Error(), "nil request")
	//})

	//t.Run("failed access token generation", func(t *testing.T) {
	//	mockAuth := new(mocks.Auth)
	//	handler := &Handler{
	//		auth: mockAuth,
	//	}
	//
	//	req := &pb.RefreshTokenRequest{
	//		RefreshToken: "invalidRefreshToken",
	//	}
	//
	//	expectedError := errors.New("failed to generate access token")
	//
	//	mockAuth.EXPECT().GenerateAccessToken(req.GetRefreshToken()).Return("", expectedError)
	//
	//	resp, err := handler.RefreshToken(context.Background(), req)
	//
	//	assert.Error(t, err)
	//	assert.Nil(t, resp)
	//	assert.Equal(t, codes.Internal, status.Code(err))
	//	assert.Contains(t, err.Error(), expectedError.Error())
	//})

	//t.Run("missing refresh token", func(t *testing.T) {
	//	mockAuth := new(mocks.Auth)
	//	handler := &Handler{
	//		auth: mockAuth,
	//	}
	//
	//	req := &pb.RefreshTokenRequest{}
	//
	//	resp, err := handler.RefreshToken(context.Background(), req)
	//
	//	assert.Error(t, err)
	//	assert.Nil(t, resp)
	//	assert.Equal(t, codes.InvalidArgument, status.Code(err))
	//	assert.Contains(t, err.Error(), "missing refresh token")
	//})
}
