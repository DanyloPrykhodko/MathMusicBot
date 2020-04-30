package store

import (
	"database/sql"
)

type Storer interface {
	DB() *sql.DB
	Dictionary() DictionaryRepositorier
}
