package main

import (
	"fmt"

	"github.com/umesshk/frenamer/internal"
)

func main() {
	err := internal.RenameFiles()

	if err != nil {
		fmt.Println("Error Occured ", err)
	}

}
