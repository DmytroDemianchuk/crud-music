package models

import "github.com/dmytrodemianchuk/crud-music/internal/domain"

type Music struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Performer   string `db:"performer"`
	RealiseYear int    `db:"realise_year"`
	Genre       string `db:"genre"`
}

func (m Music) ToDomain() domain.Music {
	return domain.Music{
		ID:          m.ID,
		Name:        m.Name,
		Performer:   m.Performer,
		RealiseYear: m.RealiseYear,
		Genre:       m.Genre,
	}
}
