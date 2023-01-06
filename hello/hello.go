package main

import (
	"fmt"
	"go_lang/greetings"
	"log"
)

func main() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	names := []string{"Mikołaj", "Natalia", "Mateusz"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
