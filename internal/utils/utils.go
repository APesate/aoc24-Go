package utils

import (
	"fmt"
	"log/slog"
	"os"
)

func ReadInput(day, part int) *os.File {
	suffix := ""

	if os.Getenv("TEST") == "1" {
		suffix = "t"
	}

	fileName := fmt.Sprintf("./Inputs/%02d-%s%d.txt", day, suffix, part)
	file, err := os.Open(fileName)
	if err != nil {
		slog.Error(err.Error())
	}

	return file
}
