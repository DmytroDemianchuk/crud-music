package service

import (
	"context"
	"github.com/dmytrodemianchuk/crud-music/internal/domain"
)

type MusicsRepository interface {
	Create(ctx context.Context, book domain.Music) error
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

func (s *Musics) Create(ctx context.Context, music domain.Music) error {
	//if music.PublishDate.IsZero() {
	//	music.PublishDate = time.Now()
	//}

	return s.repo.Create(ctx, music)
}

func (s *Musics) GetByID(ctx context.Context, id int64) (domain.Music, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Musics) GetAll(ctx context.Context) ([]domain.Music, error) {
	return s.repo.GetAll(ctx)
}

func (s *Musics) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *Musics) Update(ctx context.Context, id int64, inp domain.UpdateMusicInput) error {
	return s.repo.Update(ctx, id, inp)
}
