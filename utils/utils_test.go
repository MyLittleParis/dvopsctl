package utils

import (
    "bytes"
    "os"
    "strings"
    "testing"
)

const cmd = "dvopsctl"

func TestHomeDir(t *testing.T) {
    if HomeDir() != os.Getenv("HOME") {
        t.Errorf("The home directory is not the right one")
    }
}

func TestGetModuleRoot(t *testing.T) {
    rootDir := GetModuleRoot()
    if !strings.HasSuffix(rootDir, cmd+"/") {
        t.Errorf("%s does not contain %s, it does not seems to be the root directory of the module.", rootDir, cmd)
    } 
}

func TestCleanString(t *testing.T) {
    cases := []struct {
        str string
        name string
    }{
        {" A string with whitespace   ", "Test CleanString() with whitespace"},
        {" \"A string with double quote and whitespace\"   ", "Test CleanString() with whitespace and double quote"},
        {" 'A string with simple quote and whitespace'   ", "Test CleanString() with whitespace and simple quote"},
        {"\"A string with double quote only\"", "Test CleanString() with double quote"},
        {"'A string with simple quote only'", "Test CleanString() with simple quote"},
    }
    for _, testCase := range cases {
        t.Run(testCase.name, func(t *testing.T) {
            cleandStr := CleanString(testCase.str)

            for _, subStr := range []string{ " ", "\"", "'" } {
                if strings.HasPrefix(cleandStr, subStr) ||
                strings.HasSuffix(cleandStr, subStr) {
                    t.Errorf("String is not clean, quote or whitespace are present")
                }
            }
        })
    }
}

func TestAskForConfirmation(t *testing.T) {
    cases := []struct {
        input    string
        expected bool
        name string
    }{
        {"random\nn", false, "Test not valid response random then valid response n"},
        {"y", true, "Test valid response y"},
        {"yes", true, "Test valid response yes"},
        {"n", false, "Test valid response n"},
        {"no", false, "Test valid response no"},
        {"YES", true, "Test valid response YES"},
        {"NO", false, "Test valid response NO"},
    }

    for _, testCase := range cases {
        t.Run(testCase.name, func(t *testing.T) {
            reader := strings.NewReader(testCase.input + "\n")
            writer := new(bytes.Buffer)

            result := AskForConfirmation(testCase.name + "?", reader, writer)

            if result != testCase.expected {
                t.Errorf("AskForConfirmation(%q) = %v; want %v", testCase.input, result, testCase.expected)
            }
        })
    }
}
