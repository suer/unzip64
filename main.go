package main

import (
	"fmt"
	"log"
)

func main() {
	unzipOptions, err := parseCommandArgs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("UnzipOptions: %+v\n", unzipOptions)
}
