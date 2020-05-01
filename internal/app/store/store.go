package store

type Storer interface {
	Dictionary() DictionaryRepositorier
}
