package repository

import (
	"context"

	"github.com/lukinairina90/crud_movies/internal/domain"
	"github.com/lukinairina90/crud_movies/internal/repository/models"

	"github.com/jmoiron/sqlx"
)

type Music struct {
	db *sqlx.DB
}

func NewMusic(db *sqlx.DB) *Music {
	return &Music{db: db}
}

func (m Music) List(ctx context.Context) (domain.ListMusic, error) {
	var list []models.Music
	if err := m.db.SelectContext(ctx, &list, "SELECT * FROM music"); err != nil {
		return nil, err
	}

	dlist := make(domain.ListMusic, 0, len(list))
	for _, music := range list {
		dlist = append(dlist, music.ToDomain())
	}

	return dlist, nil
}

func (m Music) Get(ctx context.Context, id int) (domain.Music, error) {
	var music models.Music
	if err := m.db.GetContext(ctx, &music, "SELECT * FROM  music WHERE id=$1", id); err != nil {
		return domain.Music{}, err
	}

	return music.ToDomain(), nil
}

func (m Music) Create(ctx context.Context, music domain.Music) (domain.Music, error) {
	mMusic := models.Music{
		Name:        music.Name,
		Performer:   music.Performer,
		RealiseYear: music.RealiseYear,
		Genre:       music.Genre,
	}

	if err := m.db.QueryRowxContext(ctx, "INSERT INTO music (name, performer, realise_year, genre) VALUES ($1, $2, $3, $4) RETURNING *", mMusic.Name, mMusic.Performer, mMusic.ProductionYear, mMusic.Genre.Poster).StructScan(&mMusic); err != nil {
		return domain.Music{}, err
	}

	return mMusic.ToDomain(), nil
}

func (m Music) Update(ctx context.Context, id int, music domain.Music) (domain.Music, error) {
	mMusic := models.Music{
		Name:        music.Name,
		Performer:   music.Performer,
		RealiseYear: music.RealiseYear,
		Genre:       music.Genre,
	}

	if err := m.db.QueryRowxContext(ctx, "UPDATE music SET name=$1, perfomer=$2, realise_year=$3, genre=$4 WHERE id=$7 RETURNING *", mMusic.Name, mMusic.Performer, mMusic.RealiseYear, mMusic.Genre, id).StructScan(&mMusic); err != nil {
		return domain.Music{}, err
	}

	return mMusic.ToDomain(), nil
}

func (m Music) Delete(ctx context.Context, id int) error {
	if _, err := m.db.ExecContext(ctx, "DELETE FROM music WHERE id=$1", id); err != nil {
		return err
	}
	return nil
}
