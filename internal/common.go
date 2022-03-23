package internal

import "log"

func CheckError(err error, message string) {
	if err != nil {
		log.Println(message, err)
		panic(err)
	}
}
