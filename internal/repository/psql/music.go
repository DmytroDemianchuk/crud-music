package psql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/dmytrodemianchuk/crud-music/internal/domain"
)

type Musics struct {
	db *sql.DB
}

func NewMusics(db *sql.DB) *Musics {
	return &Musics{db}
}

func (r *Musics) Create(ctx context.Context, music domain.Music) error {
	_, err := r.db.Exec("INSERT INTO musics (name, artist, publish_date, genre) values ($1, $2, $3, $4)",
		music.Name, music.Artist, music.PublishDate, music.Genre)

	return err
}

func (r *Musics) GetByID(ctx context.Context, id int64) (domain.Music, error) {
	var music domain.Music
	err := r.db.QueryRow("SELECT id, name, artist, publish_date, genre FROM musics WHERE id=$1", id).
		Scan(&music.ID, &music.Name, &music.Artist, &music.PublishDate, &music.Genre)
	if err == sql.ErrNoRows {
		return music, domain.ErrMusicNotFound
	}

	return music, err
}

func (r *Musics) GetAll(ctx context.Context) ([]domain.Music, error) {
	rows, err := r.db.Query("SELECT id, name, artist, publish_date, genre FROM musics")
	if err != nil {
		return nil, err
	}

	musics := make([]domain.Music, 0)
	for rows.Next() {
		var music domain.Music
		if err := rows.Scan(&music.ID, &music.Name, &music.Artist, &music.PublishDate, &music.Genre); err != nil {
			return nil, err
		}

		musics = append(musics, music)
	}

	return musics, rows.Err()
}

func (r *Musics) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec("DELETE FROM musics WHERE id=$1", id)

	return err
}

func (r *Musics) Update(ctx context.Context, id int64, inp domain.UpdateMusicInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if inp.Name != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *inp.Name)
		argId++
	}

	if inp.Artist != nil {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argId))
		args = append(args, *inp.Artist)
		argId++
	}

	if inp.PublishDate != nil {
		setValues = append(setValues, fmt.Sprintf("publish_date=$%d", argId))
		args = append(args, *inp.Artist)
		argId++
	}

	if inp.Genre != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *inp.Genre)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE musics SET %s WHERE id=%d", setQuery, argId+1)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}
