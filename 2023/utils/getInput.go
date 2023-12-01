package utils

import (
	"fmt"
	"os"
)

func GetInput(day string) string {
	path := fmt.Sprintf("./inputs/%s.txt", day)
	data, err := os.ReadFile(path)
	if err != nil {
		return "Error reading file"
	}
	return string(data)
}
