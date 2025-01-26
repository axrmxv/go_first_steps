package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Задайте свойства предопределенного Logger, включая
	// префикс записи журнала и флаг для отключения печати
	// времени, исходного файла и номера строки.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Список имен.
	names := []string{"Alex", "Ivan", "Anna"}

	// Запрос приветственных сообщений для имен.
	messages, err := greetings.Hellos(names)
	// Если была возвращена ошибка, вывести ее на консоль и
	// выйти из программы.
	if err != nil {
		log.Fatal(err)
	}

	// Если ошибок не было, вывести возвращенную map
	// сообщений на консоль.
	fmt.Println(messages)
}
