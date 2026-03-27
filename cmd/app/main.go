package main

import (
	"fmt"
	"log"
	"merch/internal/config"
	"merch/internal/handler"
	"merch/internal/repository"
	"merch/internal/service"
	"merch/pkg/database/postgres"
	"net/http"
)

func main() {
	cfg := config.MustLoad()

	db, err := postgres.New(cfg.Database)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	log.Println("Успешно подключились к базе данных!")

	rep := repository.NewRepository(db)
	users, err := rep.GetAllUsers()

	if err != nil {
		log.Fatalf("Не вывелись пользователи: %v", err)
	}
	for _, v := range users {
		fmt.Printf("%+v\n", v)
	}

	log.Println("Нет больше юзеров")

	services := service.NewService(rep)
	handlers := handler.NewHandler(services)
	mux := handlers.InitRoutes()

	port := ":8080"
	log.Printf("Сервер запущен и слушает порт %s...", port)

	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Ошибка при работе сервера: %v", err)
	}

}
