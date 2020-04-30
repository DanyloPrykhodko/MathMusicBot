package store

import (
	"database/sql"
	"github.com/drprykhodko/MathMusicBot/internal/app/model"
)

type DictionaryRepositorier interface {
	GetKeys(*sql.DB) ([]string, error)
	Get(*sql.DB, string) (*model.Dictionary, error)
	Set(*sql.DB, *model.Dictionary) error
	Delete(*sql.DB, string) error
}
