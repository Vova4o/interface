package services

// Servicer это структура сервисного слоя
// в него прикручиваем слой хранилища!
type Servicer struct {
	db DBInterface
}

// DBInterface это интерфейсный метод слоя хранилища
// в котором указываем все методы с которыми будем работать
type DBInterface interface {
	DBUpdate(string) string
}

// NewService создает новый экземпляр сервисного слоя который в нашем случае принимает слой хранилища
func NewService(db DBInterface) *Servicer {
	return &Servicer{db: db}
}
