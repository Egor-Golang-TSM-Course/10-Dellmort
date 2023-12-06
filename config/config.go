package config

import (
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	INFO = iota + 1
	WARNING
	ERROR
)

type Config struct {
	Path           string
	Level          int8
	Flag           bool
	ReportFilePath string
	Mode           int
}

func New() (*Config, error) {
	path := flag.String("path", "", "Путь к файлу")
	level := flag.Int("level", 0, "уровень детализации анализа")
	resultFlag := flag.Bool("fs", false, "output the result to a file")
	reportFile := flag.String("report", "", "путь к файлу отчета")
	// default stdin read
	// Если стоит 1 читает из файла
	ModeFile := flag.Int("mode", 0, "mode in stdin read or file")

	flag.Parse()

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	if *path == "" {
		*path = os.Getenv("PATH_FILE")
	}
	if *level == 0 {
		levelString := os.Getenv("LOG_LEVEL")
		*level, _ = strconv.Atoi(levelString)
		if *level <= 0 || *level > 3 {
			*level = 1
		}
	}
	if *reportFile == "" {
		*reportFile = os.Getenv("REPORT_FILE")
	}

	return &Config{
		Path:           *path,
		Level:          int8(*level),
		Flag:           *resultFlag,
		Mode:           *ModeFile,
		ReportFilePath: *reportFile,
	}, nil
}
