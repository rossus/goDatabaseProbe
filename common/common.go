package common

import "log"

func CheckError(log_prefix string, err error) {
	if err != nil {
		log.Fatal(log_prefix, err)
	}
}
