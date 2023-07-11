package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacío")
	}
	return fmt.Sprintf(randomformat(), name), nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		messagge, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = messagge
	}

	return messages, nil
}

func randomformat() string {
	formatrs := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Que bueno verte, %v!",
		"¡Saludo, %v! ¡Encantado de conocerte!",
	}

	return formatrs[rand.Intn(len(formatrs))]
}
