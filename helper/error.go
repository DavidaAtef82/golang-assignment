package helper

import (
	"fmt"
	"time"
)

func CheckError(err error) bool {
	if err != nil {
		fmt.Println(err.Error(), time.Now())
		return true
	}
	return false
}
