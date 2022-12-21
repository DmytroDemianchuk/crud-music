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

func (m *Musics) Create(ctx context.Context, music domain.Music) error {
	_, err := m.db.Exec("INSERT INTO musics (name, performer, realise_year, genre) values ($1, $2, $3, $4)",
		music.Name, music.Performer, music.RealiseYear, music.Genre)

	return err
}

func (m *Musics) GetByID(ctx context.Context, id int64) (domain.Music, error) {
	var music domain.Music
	err := m.db.QueryRow("SELECT id, name, performer, realise_year, genre FROM musics WHERE id=$1", id).
		Scan(&music.ID, &music.Name, &music.Performer, &music.RealiseYear, &music.Genre)
	if err == sql.ErrNoRows {
		return music, domain.ErrMusicNotFound
	}

	return music, err
}

func (m *Musics) GetAll(ctx context.Context) ([]domain.Music, error) {
	rows, err := m.db.Query("SELECT id, name, performer, realise_year, genre FROM musics")
	if err != nil {
		return nil, err
	}

	musics := make([]domain.Music, 0)
	for rows.Next() {
		var music domain.Music
		if err := rows.Scan(&music.ID, &music.Name, &music.Performer, &music.RealiseYear, &music.Genre); err != nil {
			return nil, err
		}

		musics = append(musics, music)
	}

	return musics, rows.Err()
}

func (m *Musics) Delete(ctx context.Context, id int64) error {
	_, err := m.db.Exec("DELETE FROM musics WHERE id=$1", id)

	return err
}

func (m *Musics) Update(ctx context.Context, id int64, inp domain.UpdateMusicInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if inp.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *inp.Name)
		argId++
	}

	if inp.Performer != nil {
		setValues = append(setValues, fmt.Sprintf("performer=$%d", argId))
		args = append(args, *inp.Performer)
		argId++
	}

	if inp.RealiseYear != nil {
		setValues = append(setValues, fmt.Sprintf("realise_year=$%d", argId))
		args = append(args, *inp.RealiseYear)
		argId++
	}

	if inp.Genre != nil {
		setValues = append(setValues, fmt.Sprintf("genre=$%d", argId))
		args = append(args, *inp.Genre)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE musics SET %s WHERE id=%d", setQuery, argId+1)
	args = append(args, id)

	_, err := m.db.Exec(query, args...)
	return err
}
