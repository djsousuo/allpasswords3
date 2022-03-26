package utils

import (
	"fmt"
	"time"
)

func GetDate() string {
	year2 := time.Now().Format("2006")
	month2 := time.Now().Format("01")
	day2 := time.Now().Format("02")
	hour := time.Now().Format("15")
	min := time.Now().Format("04")
	second := time.Now().Format("05")
	return fmt.Sprintf("%s%s%s%s%s%s", year2, month2, day2, hour, min, second)
}
