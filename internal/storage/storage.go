package storage

// Storage это метод хранилища который принимает запрос
func (s *Storage) DBUpdate(text string) string {
	return text + " from storage"
}
