# Greetings

Este es un proyecto de prueba con el que estoy aprendiendo a subir y utilizar un modulo de golang. Su función es simple, recibe un nombre o un slice de nombres y la funcion retorna un saludo o un mapa de saludos respectivamente.

## Instalación

```bash
go get -u github.com/cristianino/greetings

```

## Uso


```go
package main

import (
	"fmt"
	"log"

	"github.com/cristianino/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"Cristian",
		"Alex",
		"Juan",
	}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("messages: %v\n", messages)

	mess, err := greetings.Hello("Cristian Niño")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mess)
}


```