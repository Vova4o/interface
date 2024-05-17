package storage

import "database/sql"

// Storage это структура хранилища
// тут мы храним соединение с базой данных
// можно добавить другие методы для храннеия данных и обработки
type Storage struct {
	db *sql.DB
}

// NewStorage создает новый экземпляр хранилища который в нашем случае принимает соединение с базой данных
func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}
