package usecase_test

import (
	"gitlab.golang-school.ru/potok-2/lessons/lesson-2/modules/internal/usecase"
	"testing"
)

func TestHello(t *testing.T) {
	hello := usecase.NewHello()

	actual := hello.Say()
	expected := "Hello!"

	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
