package bot

import (
	"database/sql"
	"errors"
	"github.com/drprykhodko/MathMusicBot/internal/app/store/sqlstore"
	_ "github.com/lib/pq"
	"sync"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	s := sqlstore.New(db)

	b := newBot(config.Token, s)

	wg := sync.WaitGroup{}
	wg.Add(1)

	b.launch()

	wg.Wait()

	return errors.New("bot stopped")
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
