package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

const noArgErrMsg = "No command provided"
const unknownCmd = "Unknown command"


func Run(args []string) (int, error) {
    var rootError error
    var errCode int

    serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
    serverOpen := serverCmd.Bool("open", false, "Open local SERVER_NAME if found in .env file")
    serverCmd.Usage = func() {
        fmt.Println("server [options]")
        serverCmd.PrintDefaults()
        fmt.Println("")
    }

    // Commands list to use it in usage function
    commands := []*flag.FlagSet{
            serverCmd,
    }

    var usage func()
    usage = func ()  {
        
        fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", args[0])

        for _, cmd := range commands {
            cmd.Usage()
        }
    }
   
    if len(args) == 1 {
        usage()
        return -1, errors.New(noArgErrMsg)
    }

    switch args[1] {
    case "server":
        serverCmd.Parse(args[2:])
        if *serverOpen {
            errCode, rootError = ServerOpen()
        }
        default: 
            fmt.Println("Unknown command: " + args[1])
            usage()
            return -1, errors.New(unknownCmd)
    }
    
    return errCode, rootError
}
