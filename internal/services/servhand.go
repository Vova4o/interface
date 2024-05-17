package services

// DoSomething метод сервисного слоя который обрабатывает запрос
// и передает его на слой хранилища!
func (s *Servicer) DoSomething(text string) string {
	result := s.db.DBUpdate(text + " from service ")
	return result
}
