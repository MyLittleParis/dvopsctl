package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MyLittleParis/dvopsctl/utils"
)

const EnvFile = ".env"
const EnvPrefix = "DVOPS_"

func GetEnvFilePath() string {
    return utils.GetModuleRoot() + EnvFile
}

func SanitizeName(name string) string {
    name = strings.ReplaceAll(name, " ", "_")
    name = strings.ToUpper(name)

    return name
}

func setEnvVar(variable, value string) {
    value = utils.CleanString(value)

    os.Setenv(EnvPrefix + variable, value)
}

func SetEnvVars() {
    rootPath := utils.GetModuleRoot()
    if content, err := os.Open(rootPath + EnvFile); err == nil {
        scanner := bufio.NewScanner(content)
        for scanner.Scan() {
            if variable, value, found := strings.Cut(scanner.Text(), "="); found {
                setEnvVar(variable, value)
            }
        }
    }
}

func ListEnvVars() {
    fmt.Println("List environment variables")
    for _, env := range os.Environ() {
        variable, value, _ := strings.Cut(env, "=")
        if strings.Contains(variable, EnvPrefix) {
            fmt.Println(variable, value)
        }
    }
}
