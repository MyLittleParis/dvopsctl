package utils

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "os/exec"
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

// Git functions
func GitClone(repository string) {
    git := exec.Command("git", "clone", repository)
    err := git.Start()

    log.Printf("git clone %s", repository)

	if err != nil {
		log.Fatal(err)
	}

	err = git.Wait()

    if err != nil {
        log.Printf("Command finished with error: %v", err)
    }
}

func GitPull() {
    git := exec.Command("git", "pull")
    err := git.Start()

    log.Printf("git pull")

	if err != nil {
		log.Fatal(err)
	}

	err = git.Wait()

    if err != nil {
        log.Printf("Command finished with error: %v", err)
    }
}

func GitCheckout(branch string) {
    git := exec.Command("git", "checkout", branch)
    err := git.Start()

    log.Printf("git checkout %s", branch)

	if err != nil {
		log.Fatal(err)
	}

	err = git.Wait()

    if err != nil {
        log.Printf("Command finished with error: %v", err)
    }
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
