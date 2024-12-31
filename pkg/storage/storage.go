package storage

import "database/sql"

type DataStorage struct {
	*sql.DB
}

func NewDataStorage() *DataStorage {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}

	return &DataStorage{db}
}

func (ds *DataStorage) init() {
	query := `CREATE TABLE IF NOT EXISTS campaigns (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL unique,
		email TEXT
	)`
	_, err := ds.Exec(query)
	if err != nil {
		panic(err)
	}
}

func (ds *DataStorage) Close() {
	ds.DB.Close()
}
