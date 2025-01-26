package greetings

import (
	"fmt"

	"errors"

	"math/rand"
)

// Hello возвращает приветствие для указанного человека.
func Hello(name string) (string, error) {
	// Если имя не указано, вернуть ошибку сообщением.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Если имя получено, вернуть значение, которое встраивает имя
	// в приветственное сообщение.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos возвращает map, которая связывает каждого из указанных людей
// с приветственным сообщением.
func Hellos(names []string) (map[string]string, error) {
	// map для связывания имен с сообщениями.
	messages := make(map[string]string)

	for _, name := range names {
		// Проходим по полученному фрагменту имен, вызывая
		// функцию Hello, чтобы получить сообщение для каждого имени.
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// В map связываем полученное сообщение с
		// именем.
		messages[name] = message
	}
	return messages, nil
}

// randomFormat возвращает одно из набора приветственных сообщений. Возвращаемое сообщение
// выбирается случайным образом.
func randomFormat() string {
	formats := []string{
		"Hi %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Возвращает случайно выбранный формат сообщения, указывая
	// случайный индекс для среза форматов.
	return formats[rand.Intn(len(formats))]
}
