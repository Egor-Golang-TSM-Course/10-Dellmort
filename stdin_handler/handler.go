package stdinhandler

import (
	"bufio"
	"os"
)

type Handler struct {
	body []byte
}

func New(b []byte) *Handler {
	return &Handler{
		body: b,
	}
}

func (h *Handler) Read(b []byte) (n int, err error) {
	for {
		reader := bufio.NewReader(os.Stdin)
		return reader.Read(b)
	}
}
