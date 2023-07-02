package errors

import err "errors"

var (
	ErrUnauthorized      = err.New("unauthorized")
	ErrUserNotFound      = err.New("user not found")
	ErrMemberNotFound    = err.New("member not found")
	ErrWorkspaceNotFound = err.New("workspace not found")
)
