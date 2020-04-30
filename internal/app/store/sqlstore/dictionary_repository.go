package sqlstore

import (
	"database/sql"
	"errors"
	"github.com/drprykhodko/MathMusicBot/internal/app/model"
	"github.com/drprykhodko/MathMusicBot/internal/app/store"
)

var (
	ErrorUnknownKey      = errors.New("unknown key")
	ErrorNothingToDelete = errors.New("nothing to delete")
)

type DictionaryRepository struct {
	store store.Storer
}

func (d DictionaryRepository) GetKeys(db *sql.DB) (keys []string, err error) {
	rows, err := db.Query(
		`SELECT "key" FROM "dictionary" ORDER BY "key"`,
	)
	if err != nil {
		return
	}

	for rows.Next() {
		var key string
		err = rows.Scan(&key)
		if err != nil {
			return nil, err
		}
		keys = append(keys, key)
	}

	return
}

func (d DictionaryRepository) Get(db *sql.DB, key string) (dictionary *model.Dictionary, err error) {
	dictionary = &model.Dictionary{}
	err = db.QueryRow(
		`SELECT "key", "value" FROM "dictionary" WHERE LOWER("key") = LOWER($1)`,
		key,
	).Scan(
		&dictionary.Key,
		&dictionary.Value,
	)
	if err == sql.ErrNoRows {
		return nil, ErrorUnknownKey
	}

	return
}

func (d DictionaryRepository) Set(db *sql.DB, dictionary *model.Dictionary) (err error) {
	_, err = db.Exec(
		`INSERT INTO "dictionary"("key", "value") VALUES ($1, $2)`,
		dictionary.Key,
		dictionary.Value,
	)

	return
}

func (d DictionaryRepository) Delete(db *sql.DB, key string) (err error) {
	result, err := db.Exec(
		`DELETE FROM "dictionary" WHERE LOWER("key") = LOWER($1)`,
		key,
	)
	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrorNothingToDelete
	}

	return
}
