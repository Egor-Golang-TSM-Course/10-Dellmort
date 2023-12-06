package stdinhandler

import (
	"bufio"
	"os"
	"strings"
)

type Handler struct {
	statsTypeMap map[string]int
}

func New() *Handler {
	return &Handler{
		statsTypeMap: map[string]int{
			"INFO":    0,
			"WARNING": 0,
			"ERROR":   0,
		},
	}
}

func (h *Handler) ReadConsole() (map[string]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		if strings.Contains(text, "INFO") {
			h.statsTypeMap["INFO"]++
		}
		if strings.Contains(text, "WARNING") {
			h.statsTypeMap["WARNING"]++
		}
		if strings.Contains(text, "ERROR") {
			h.statsTypeMap["ERROR"]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return h.statsTypeMap, nil
}
