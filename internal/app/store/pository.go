package store

import (
	"github.com/drprykhodko/MathMusicBot/internal/app/model"
)

type DictionaryRepositorier interface {
	GetKeys() ([]string, error)
	Get(string) (*model.Dictionary, error)
	Set(*model.Dictionary) error
	Delete(string) error
}
