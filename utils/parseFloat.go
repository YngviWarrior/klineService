package utils

import (
	"log"
	"strconv"
)

func ParseFloat(s string) (f float64) {
	f, err := strconv.ParseFloat(s, 64)

	if err != nil {
		log.Panicln("PF 01: ", err)
	}

	return
}
