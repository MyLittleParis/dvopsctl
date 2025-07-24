package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/MyLittleParis/dvopsctl/config"
	"github.com/MyLittleParis/dvopsctl/utils"
)

var envFilePath = config.GetEnvFilePath()

func AddEnvVar(variable, value string) (int, error) {
    value = utils.CleanString(value)
    variable = config.SanitizeName(variable)

    if oldValue, found := os.LookupEnv(config.EnvPrefix + variable); found {
        fmt.Printf("%s already exist, value: %s\n", variable, oldValue)

        if utils.AskForConfirmation("Do you want to update value ?", os.Stdin, os.Stdout) {
            _, err := UpdateEnvVar(value, variable)

            if err != nil {
                return -1, err
            }

            return 0, errors.New("Variable updated")
        } else {
            return 0, errors.New("Variable already set")
        }
    }

    newLine := fmt.Sprintf("%s=%s", variable, value)
    envFile, err := os.OpenFile(envFilePath, os.O_APPEND|os.O_WRONLY, 0644)

    if err != nil {
        fmt.Println(err)
    }

    defer envFile.Close()

    if _, err = envFile.Write([]byte("\n"+newLine)); err != nil {
        fmt.Println(err)
    }

    return 0, nil
}


func ListEnvVars() (int, error) {

    config.ListEnvVars()

    return -1, errors.New(noEnv)
}

func SetEnvVars() (int, error) {

    config.SetEnvVars()

    return -1, errors.New(noEnv)
}

func UpdateEnvVar(variable, value string) (int, error) {
    value = utils.CleanString(value)
    variable = config.SanitizeName(variable)

    pattern := fmt.Sprintf("s/%s=.*/%s=%s/g", variable, variable, value)

    sed := exec.Command("sed", "-i", "-e", pattern, envFilePath)
    if err := sed.Start(); err != nil {
        return -1, err
    }

    return 0, nil
}
