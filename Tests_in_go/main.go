package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100.0, 10.0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("result of division is", result)
}

func divide(a, b float32) (float32, error) {
	var result float32

	if b == 0 {
		return result, errors.New("Cannot divide by 0")
	}

	result = a / b
	return result, nil
}