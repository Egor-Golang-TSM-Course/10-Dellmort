package processor

import (
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

	reportFile, err := os.OpenFile(config.ReportFilePath, os.O_CREATE|os.O_RDWR, 0744)
	if err != nil {
		return nil, err
	}

	return &analyzer{
		file:       file,
		level:      config.Level,
		modeRead:   config.Mode,
		reportFile: reportFile,
		handler:    handler.New(),
		flag:       config.Flag,
	}, nil
}

func (a *analyzer) Analysis() error {
	defer a.reportFile.Close()

	logType, err := a.analysis()
	if err != nil {
		return err
	}

	var (
		info      = logType["INFO"]
		warning   = logType["WARNING"]
		errorType = logType["ERROR"]
	)

	if a.reportFile != nil && a.flag {
		switch a.level {
		case 1:
			_, err := a.reportFile.WriteString(
				fmt.Sprintf("INFO: %d\nWARNING:%d\nERROR:%d", info, warning, errorType),
			)
			if err != nil {
				return err
			}
		case 2:
			_, err := a.reportFile.WriteString(
				fmt.Sprintf("WARNING:%d\n ERROR:%d", warning, errorType),
			)
			if err != nil {
				return err
			}
		case 3:
			_, err := a.reportFile.WriteString(
				fmt.Sprintf("ERROR:%d", errorType),
			)
			if err != nil {
				return err
			}
		}
	} else {
		fmt.Printf("INFO:%d, WARNING:%d, ERROR:%d", info, warning, errorType)
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
		return a.handler.ReadConsole()
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
