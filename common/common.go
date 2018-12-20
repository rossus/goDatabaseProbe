package common

import "log"

const (
	FATAL = "[FATAL] "
	ERROR = "[ERROR] "
)

func CheckError(log_prefix string, err error) {
	if err != nil {
		log.Fatal(FATAL, log_prefix, err)
	}
}

//Reserved for non-fatal errors
func WhatDoesItSay(log_prefix string, err error)  {
	if err != nil {
		log.Print(ERROR, log_prefix, err, "\n")
	}
}