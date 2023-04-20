package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

func main() {
	// Чтение файла с Swagger-схемой
	swaggerBytes, err := ioutil.ReadFile("./swagger/swagger.json")
	if err != nil {
		log.Fatal("Ошибка чтения файла с Swagger-схемой: ", err)
	}

	// Создание middleware из Swagger-спецификации
	handlerMiddleware := middleware.Spec("/docs", swaggerBytes, nil)

	// Запуск HTTP-сервера на порту 8080
	log.Fatal(http.ListenAndServe(":8080", handlerMiddleware))
}
