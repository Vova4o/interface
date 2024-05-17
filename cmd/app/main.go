package main

import (
	"database/sql"
	"fmt"

	"github.com/Vova4o/interface/internal/handlers"
	"github.com/Vova4o/interface/internal/services"
	"github.com/Vova4o/interface/internal/storage"
)

func main() {
	// чисто теоретический пример
	// впредставим что создаем соединение с базой данных
	db := &sql.DB{}

	// соединеие с базой данных передаем в хранилище
	stor := storage.NewStorage(db)
	// хранилище передаем на сервисный слой
	serv := services.NewService(stor)
	// сервисный слой передаем на обработчик
	hand := handlers.NewHandler(serv)

	// вызываем метод обработчика
	answer := hand.Handle("Hello World!")

	// выводим результат
	fmt.Println(answer)
}
