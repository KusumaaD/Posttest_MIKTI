package utils

import "log"

// CheckError handles the error and logs it
func CheckError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}