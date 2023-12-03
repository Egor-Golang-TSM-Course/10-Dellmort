package processor

import (
	"encoding/json"
	"fmt"
	"io"
	"lesson10/config"
	handler "lesson10/stdin_handler"
	"os"
	"strings"
)

type analyzer struct {
	file       *os.File
	level      int8
	handler    *handler.Handler
	modeRead   int
	reportFile *os.File
	flag       bool
}

func New(config *config.Config, b []byte) (*analyzer, error) {
	file, err := os.OpenFile(config.Path, os.O_RDONLY, 6544)
	if err != nil {
		return nil, err
	}
	reportFile, err := os.OpenFile(config.ReportFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0744)
	if err != nil {
		return nil, err
	}
	return &analyzer{
		file:       file,
		level:      config.Level,
		modeRead:   config.Mode,
		reportFile: reportFile,
		handler:    handler.New(b),
		flag:       config.Flag,
	}, nil
}

func (a *analyzer) Analysis() error {
	defer a.reportFile.Close()

	logType, err := a.analysis()
	if err != nil {
		return err
	}

	body, err := json.MarshalIndent(logType, " ", "  ")
	if err != nil {
		return err
	}

	if a.reportFile != nil && a.flag {
		a.reportFile.Write(body)
	} else {
		fmt.Println(logType)
	}
	return nil
}

func (a *analyzer) analysis() (map[string]int, error) {
	logType := map[string]int{
		"INFO":    0,
		"WARNING": 0,
		"ERROR":   0,
	}
	var (
		buffer []byte
		err    error
	)
	if a.modeRead == 0 {
		buffer, err = io.ReadAll(a.handler)
		if err != nil {
			return nil, err
		}
	} else {
		buffer, err = io.ReadAll(a.file)
		if err != nil {
			return nil, err
		}
	}
	defer a.file.Close()

	fields := strings.Fields(string(buffer))
	for _, value := range fields {
		field := strings.Trim(value, "[]")
		if _, ok := logType[field]; ok {
			logType[field]++
		}
	}

	return logType, nil
}
