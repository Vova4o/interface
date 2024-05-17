package handlers

// Handler это метод обработчика который принимает запрос
// и передает его на сервисный слой!
func (s *Handler) Handle(text string) string {
	textedit := s.s.DoSomething(text + " from handler ")
	return textedit
}
