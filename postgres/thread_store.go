package postgres

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewThreadStore(db *sqlx.DB) *ThreadStore {
	return &ThreadStore{
		DB: db,
	}
}

type ThreadStore struct {
	*sqlx.DB
}

func (s *ThreadStore) Thread(id uuid.UUID) (goreddit.Thread, error) {
	panic("not implemented") // TODO: Implement
}

func (s *ThreadStore) Threads() ([]goreddit.Thread, error) {
	panic("not implemented") // TODO: Implement
}

func (s *ThreadStore) UpdateThread(t *goreddit.Thread) error {
	panic("not implemented") // TODO: Implement
}

func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}
