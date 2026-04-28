package main

import (
	"fmt"
	"github.com/google/uuid"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-2/modules/internal/usecase"
	_ "go.uber.org/automaxprocs"
)

func main() {
	hello := usecase.NewHello()
	fmt.Println(hello.Say())

	fmt.Println("UUID:", uuid.New().String())
}
