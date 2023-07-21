package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func Open(config Config) (*Store, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	store := NewStore(db)
	return store, nil
}

func (store *Store) CheckStoreConnection() bool {
	return store.db.Ping() != nil
}

func (store *Store) Close() error {
	return store.db.Close()
}
