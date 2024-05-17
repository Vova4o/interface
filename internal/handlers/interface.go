package handlers

// Handler это метод обработчика который принимает запрос
// в него прикручиваем сервисный слой!
type Handler struct {
	s Servicer
}

// Servicer это интерфейсный метод сервисного слоя
type Servicer interface {
	DoSomething(string) string
}

// NewHandler создает новый экземпляр обработчика который в нашем случае принимает сервисный слой
// и возвращает его
func NewHandler(s Servicer) *Handler {
	return &Handler{s: s}
}
