package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/MyLittleParis/dvopsctl/utils"
	"github.com/pkg/browser"
)

const noEnv = "no env file"

var envFiles = []string{".env", ".docker/.env"}

func ServerOpen() (int, error) {
    servername := searchInEnvFile("SERVER_NAME")
    if value, found := strings.CutPrefix(servername, "${COMPOSE_PROJECT_NAME}"); found {
        project := searchInEnvFile("COMPOSE_PROJECT_NAME")
        servername = project+value
    }

    if servername != "" {
        url := "https://" + servername

        fmt.Println("Opening " + url)
        if servername != "dvopsctl.docker.localhost" {
            browser.OpenURL(url)
        }

        return 0, nil
    }

    fmt.Println("No SERVER_NAME found in .env or .docker/.env files.")
    fmt.Println("Or no env file found.")

    return -1, errors.New(noEnv)
}

func searchInEnvFile(envVar string) string {
    for _, file := range envFiles {
        if content, err := os.Open(file); err == nil {
            scanner := bufio.NewScanner(content)
            for scanner.Scan() {
                if value := searchEnvVar(envVar, scanner.Text()); value != "" {
                    return utils.RemoveQuote(value)
                }
            }
        }
    }

    return ""
}

func searchEnvVar(envvar, line string) string {
    if value, found := strings.CutPrefix(line, envvar+"="); found {
        return value
    }
    return ""
}
