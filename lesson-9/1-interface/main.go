package main

import (
	"fmt"

	pkg_file "gitlab.golang-school.ru/potok-2/lessons/lesson-9/1-interface/pkg/file"
	"gitlab.golang-school.ru/potok-2/lessons/lesson-9/1-interface/pkg/stdout"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func worker(w Writer, msg string) {
	n, err := w.Write([]byte(msg))
	if err != nil {
		fmt.Println("worker: write error:", err)
	}

	fmt.Printf("worker: success write %d bytes\n", n)
}

func main() {
	file, err := pkg_file.NewWriter("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	console := stdout.NewWriter()

	worker(file, "Hello, file!\n")
	worker(console, "Hello, stdout!\n")
}
