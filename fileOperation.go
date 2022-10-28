package fileOperation

import (
	"os"

	"github.com/golang-module/carbon/v2"
)

func WriteDateToFile(date string) {
	d1 := []byte(date)
	err := os.WriteFile("date", d1, 0644)
	check(err)
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

func MonthCompare(month string) string {
	var result string

	switch month {
	case "Ocak":
		result = "1"
	case "Şubat":
		result = "2"
	case "Mart":
		result = "3"
	case "Nisan":
		result = "4"
	case "Mayıs":
		result = "5"
	case "Haziran":
		result = "6"
	case "Temmuz":
		result = "7"
	case "Ağustos":
		result = "8"
	case "Eylül":
		result = "9"
	case "Ekim":
		result = "10"
	case "Kasım":
		result = "11"
	case "Aralık":
		result = "12"
	}

	return result
}
