package common

import "log"

//LogFatal checks	 fatal error
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
