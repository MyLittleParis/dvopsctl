package utils

import (
    "os"
    "fmt"
    "strings"
)

func HomeDir() string {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return homeDir
}

func RemoveQuote(s string) string {
    s = strings.ReplaceAll(s, "\"", "")
    s = strings.ReplaceAll(s, "'", "")
    return s
}
