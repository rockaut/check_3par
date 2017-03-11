package main

import (
	"fmt"
	"os"

	"github.com/rockaut/check_3par/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
