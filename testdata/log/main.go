package main

import "log"

func main() {
	myLog("ayo")
}

func myLog(format string, args ...interface{}) {
	const prefix = "[my] "
	log.Printf(prefix+format, args...)
}
