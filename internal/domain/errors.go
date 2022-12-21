package domain

import "errors"

var (
	ErrMusicNotFound       = errors.New("music not found")
	ErrRefreshTokenExpired = errors.New("refresh token expired")
)
