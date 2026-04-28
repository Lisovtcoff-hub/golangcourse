package stdout

import "fmt"

type Writer struct{}

func NewWriter() *Writer {
	return &Writer{}
}

func (s *Writer) Write(p []byte) (int, error) {
	return fmt.Println(string(p))
}
