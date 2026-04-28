package file

import "os"

type Writer struct {
	file *os.File
}

func NewWriter(filename string) (*Writer, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return &Writer{file: f}, nil
}

func (w *Writer) Write(p []byte) (int, error) {
	n, err := w.file.Write(p)
	if err != nil {
		return n, err
	}

	return n, nil
}

func (w *Writer) Close() error {
	return w.file.Close()
}
