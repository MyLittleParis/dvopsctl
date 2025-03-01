package main

import (
	"flag"
	"fmt"
	"os"

	server "github.com/MyLittleParis/dvopsctl/commands"
)

func main() {
    serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
    serverOpen := serverCmd.Bool("open", false, "Open local SERVER_NAME if found in .env file")

    if len(os.Args) < 2 {
        fmt.Println("A command must be passed as argument.")
        fmt.Println("    For e.g :")
        fmt.Println("        dvopsctl server -open")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "server":
        serverCmd.Parse(os.Args[2:])
        if *serverOpen {
            server.Open()
        }
    default: 
        flag.Usage()
    }
}
