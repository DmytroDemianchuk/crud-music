package service

import (
	"context"

	"github.com/dmytrodemianchuk/crud-music/internal/domain"
)

type MusicRepository interface {
	List(ctx context.Context) (domain.ListMusic, error)
	Get(ctx context.Context, id int) (domain.Music, error)
	Create(ctx context.Context, music domain.Music) (domain.Music, error)
	Update(ctx context.Context, id int, music domain.Music) (domain.Music, error)
	Delete(ctx context.Context, id int) error
}

type Music struct {
	musicRepository MusicRepository
}

func NewMusic(musicRepository MusicRepository) *Music {
	return &Music{musicRepository: musicRepository}
}

func (m Music) List(ctx context.Context) (domain.ListMusic, error) {
	return m.musicRepository.List(ctx)
}

func (m Music) Get(ctx context.Context, id int) (domain.Music, error) {
	return m.musicRepository.Get(ctx, id)
}

func (m Music) Create(ctx context.Context, music domain.Music) (domain.Music, error) {
	return m.musicRepository.Create(ctx, music)
}

func (m Music) Update(ctx context.Context, id int, music domain.Music) (domain.Music, error) {
	return m.musicRepository.Update(ctx, id, music)
}

func (m Music) Delete(ctx context.Context, id int) error {
	return m.musicRepository.Delete(ctx, id)
}
