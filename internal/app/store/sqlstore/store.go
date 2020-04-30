package sqlstore

import (
	"database/sql"
	"github.com/drprykhodko/MathMusicBot/internal/app/store"
)

type Store struct {
	db         *sql.DB
	dictionary store.DictionaryRepositorier
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s Store) DB() *sql.DB {
	return s.db
}

func (s Store) Dictionary() store.DictionaryRepositorier {
	if s.dictionary != nil {
		return s.dictionary
	}

	s.dictionary = &DictionaryRepository{
		store: s,
	}

	return s.dictionary
}
