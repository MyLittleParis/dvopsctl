package cmd

import (
	"os"
	"strings"
	"testing"
)

const cmd = "dvopsctl"

func TestRun(t *testing.T) {
    cases := []struct {
        args []string
        errMsg string
        errorCode int
        name string
        testDir string
    }{
        {[]string{cmd}, noArgErrMsg, -1, "Test Run() without arguments", ""},
        {[]string{cmd, "unknown"}, unknownCmd, -1, "Test Run() with unknown command", ""},
        {[]string{cmd, "server", "-open"}, noEnv, -1, "Test Run() without .env file", ""},
        {[]string{cmd, "server", "-open"}, "", 0, "Test Run() with .env file", "../testdata/"},
    }

    for _, testCase := range cases {
        t.Run(testCase.name, func(t *testing.T) {
            if testCase.testDir != "" {
                err := os.Chdir(testCase.testDir)
                if err != nil {
                    t.Fatalf("Failed to change directory to testdata: %v", err)
                }
            }

            cmdWithArg := strings.Join(testCase.args, " ")

            errCode, err := Run(testCase.args)

            if errCode != testCase.errorCode {
                t.Errorf("%s error code = %d; want %d", cmdWithArg, errCode, testCase.errorCode)
            }
            if testCase.errMsg != "" {
                if err.Error() != testCase.errMsg {
                    t.Errorf("%s return %s, want %s", cmdWithArg, err.Error(), testCase.errMsg)
                }
            }else if err != nil {
                t.Errorf("%s return %s, want %s", cmdWithArg, err.Error(), "nil")
            }
        })
    }
}
