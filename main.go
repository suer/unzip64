package main

import (
	"fmt"
	"os"
)

func main() {
	unzipOptions, err := parseCommandArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = unzip(*unzipOptions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
