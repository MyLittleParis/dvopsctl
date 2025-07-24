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

    SetEnvVars()

    configCmd := flag.NewFlagSet("config", flag.ExitOnError)
    // subcommands
    configGet := configCmd.Bool("get", false, "Get dvopsctl environment variables")
    configSet := configCmd.Bool("set", false, "Set dvopsctl config from environment variable on .env file")
    configAdd := configCmd.Bool("add", false, "Add dvopsctl config, add environment variable on .env file")
    configUpdate := configCmd.Bool("update", false, "Update dvopsctl config, update environment variable on .env file")
    // args
    configName := configCmd.String("n", "", "Environment variable's name on .env file")
    configValue := configCmd.String("v", "", "Environment variable's value on .env file")

    configCmd.Usage = func() {
        fmt.Println("config [options]")
        configCmd.PrintDefaults()
        fmt.Println("")
    }

    serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
    // subcommands
    serverOpen := serverCmd.Bool("open", false, "Open local SERVER_NAME if found in .env file")

    serverCmd.Usage = func() {
        fmt.Println("server [options]")
        serverCmd.PrintDefaults()
        fmt.Println("")
    }

    // Commands list to use it in usage function
    commands := []*flag.FlagSet{
        configCmd,
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
    case "config":
        configCmd.Parse(args[2:])
        if *configGet {
            ListEnvVars()
        }
        if *configSet {
            SetEnvVars()
        }
        if *configAdd {
            if *configName != "" && *configValue != "" {
                AddEnvVar(*configName, *configValue)
            } else {
                fmt.Println("config -add command needs -n NAME and -v VALUE of your environment variable.")
            }
        }
        if *configUpdate {
            if *configName != "" && *configValue != "" {
                AddEnvVar(*configName, *configValue)
            } else {
                fmt.Println("config -add command needs -n NAME and -v VALUE of your environment variable.")
            }
        }
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
