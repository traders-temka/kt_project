package repository

import (
	"database/sql"
	"kt_project/internal/models"

	_ "modernc.org/sqlite" // driver for sqlite
)

type SqlStorage struct {
	db *sql.DB
}

func NewSqlStorage(path string) (*SqlStorage, error) {
	db, err := sql.Open("sqlite", path)

	if err != nil {
		return nil, err
	}

	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS stats (
        symbol TEXT,
        price REAL,
        source TEXT,
        timedump INTEGER
    )`)
	if err != nil {
		return nil, err
	}
	statement.Exec()

	return &SqlStorage{db: db}, nil
}

func (s *SqlStorage) Save(stat models.Stat) error {
	// query request to sql/db
	query := `INSERT INTO stats (symbol, price, source, timedump) VALUES (?, ?, ?, ?)`
	_, err := s.db.Exec(query, stat.Name, stat.Price, stat.Source, stat.Timedump)
	return err
}
