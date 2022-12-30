package service

import (
	"context"
	"time"

	"github.com/dmytrodemianchuk/crud-music/internal/domain"
)

type MusicsRepository interface {
	Create(ctx context.Context, music domain.Music) error
	GetByID(ctx context.Context, id int64) (domain.Music, error)
	GetAll(ctx context.Context) ([]domain.Music, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, inp domain.UpdateMusicInput) error
}

type Musics struct {
	repo MusicsRepository
}

func NewMusics(repo MusicsRepository) *Musics {
	return &Musics{
		repo: repo,
	}
}

func (m *Musics) Create(ctx context.Context, music domain.Music) error {
	if music.PublishDate.IsZero() {
		music.PublishDate = time.Now()
	}

	return m.repo.Create(ctx, music)
}

func (m *Musics) GetByID(ctx context.Context, id int64) (domain.Music, error) {
	return m.repo.GetByID(ctx, id)
}s

func (m *Musics) GetAll(ctx context.Context) ([]domain.Music, error) {
	return m.repo.GetAll(ctx)
}

func (m *Musics) Delete(ctx context.Context, id int64) error {
	return m.repo.Delete(ctx, id)
}

func (m *Musics) Update(ctx context.Context, id int64, inp domain.UpdateMusicInput) error {
	return m.repo.Update(ctx, id, inp)
}
