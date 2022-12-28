package domain

import (
	"errors"
	"time"
)

var (
	ErrMusicNotFound = errors.New("music not found")
)

type Music struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Artist      string    `json:"artist"`
	PublishDate time.Time `json:"publish_date"`
	Genre       int       `json:"genre"`
}

type UpdateMusicInput struct {
	Name        *string    `json:"name"`
	Artist      *string    `json:"artist"`
	PublishDate *time.Time `json:"publish_date"`
	Genre       *int       `json:"genre"`
}
