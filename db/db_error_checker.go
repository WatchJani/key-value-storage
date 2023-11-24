package db

import "fmt"

func ErrorChecker(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
