package errorhandling

import (
	"fmt"
	"log"
)

func FailIf(err error, errstr string) {
	if err != nil {
		log.Fatal(fmt.Sprintf("%s %s", errstr, err.Error()))
	}
}
