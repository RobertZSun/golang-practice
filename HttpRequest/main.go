package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	fmt.Println("Hello", reflect.TypeOf(os.Args))
	if len(os.Args) == 2 {
		fileName := os.Args[1]

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		io.Copy(os.Stdout, file)

	} else {
		os.Exit(1)
	}

}
