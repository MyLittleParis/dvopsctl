package utils

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "runtime"
    "strings"
)

// Paths functions
func HomeDir() string {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return homeDir
}

func GetModuleRoot() string {
    _, filename, _, _ := runtime.Caller(0)
    rootDir, _ := strings.CutSuffix(filepath.Dir(filename), "utils")

    return rootDir
}

// Custom strings functions
func CleanString(s string) string {
    cleanString := strings.TrimSpace(s)
    cleanString = strings.Trim(cleanString, "\"")
    cleanString = strings.Trim(cleanString, "'")

    return cleanString
}

func RemoveQuote(s string) string {
    s = strings.ReplaceAll(s, "\"", "")
    s = strings.ReplaceAll(s, "'", "")
    return s
}

// Others utils functions

func AskForConfirmation(s string, reader io.Reader, writer io.Writer) bool {
    scanner := bufio.NewScanner(reader)

    for {
        fmt.Fprintf(writer, "%s [y/n]: ", s)

        if scanner.Scan() {
            response := strings.ToLower(strings.TrimSpace(scanner.Text()))

            if response == "y" || response == "yes" {
                return true
            } else if response == "n" || response == "no" {
                return false
            }
        } else if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
    }
}
