package database

import (
	"database/sql"

	"DiscordTemplate/pkg/logger"
)

type sqliteDB struct {
	db *sql.DB
	logger logger.Logger
}

func NewSqliteDB(logger logger.Logger) *sqliteDB {
	db, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		logger.Fatal("Failed to connect to database: %v", err)
	}
	db.SetMaxOpenConns(1)
	sqlitedb := &sqliteDB{
		db: db,
		logger: logger,
	}
	err = sqlitedb.InitDB()
	if err != nil {
		logger.Warn("Failed to initialize database: %v", err)
	}
	return sqlitedb
}

func (s *sqliteDB) InitDB() error {
	query := "CREATE TABLE IF NOT EXISTS userdata (userID TEXT, secret TEXT)"
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqliteDB) Close() {
	err := s.db.Close()
	if err != nil {
		s.logger.Error("Failed to close database connection")
	}
}

func (s *sqliteDB) Write(userID string, secret string) error {
	query := "INSERT INTO userdata (userID, secret) VALUES (?, ?)"
	_, err := s.db.Exec(query, userID, secret)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqliteDB) Remove(userID string) error {
	query := "DELETE FROM userdata WHERE userID = ?"
	_, err := s.db.Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}
