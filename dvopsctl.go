package main

import (
	"fmt"
	"os"

	cmd "github.com/MyLittleParis/dvopsctl/commands"
)

func main() {
    if exitCode, err := cmd.Run(os.Args); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(exitCode)
    }
}
