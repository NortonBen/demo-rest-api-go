package util

import (
	"log"
	"os"
	"strconv"
)

func LookupEnvInt(name string) (int64, bool) {
	val, exit := os.LookupEnv(name)
	if !exit {
		return 0, false
	}
	number, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Println("error convert Data", err)
		return 0, false
	}
	return number, true
}
