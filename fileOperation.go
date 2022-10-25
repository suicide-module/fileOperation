package fileOperation

import (
	"os"

	"github.com/golang-module/carbon/v2"
)

func WriteFromFile(date string) {
	d1 := []byte(date)
	err := os.WriteFile("date", d1, 0644)
	check(err)
}

func ReadFromFile() string {
	result, err := os.ReadFile("date")
	check(err)
	return string(result)
}

func ReadDateFromFile() carbon.Carbon {
	result, err := os.ReadFile("date")
	check(err)
	return carbon.Parse(string(result))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
