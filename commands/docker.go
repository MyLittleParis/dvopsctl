package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/MyLittleParis/dvopsctl/utils"
)

const dockerfile = "Dockerfile"
const defaultRegistry = "ghcr.io/mylittleparis"
const defaultGithub = "https://github.com/MyLittleParis" 
const defaultRepositoryName = "devops-docker-php"

var defaultRegistryRepo = fmt.Sprintf("%s/%s", defaultRegistry, defaultRepositoryName)
var defaultRepository = fmt.Sprintf("%s/%s", defaultGithub, defaultRepositoryName)
var defaultPath = []string {"./", ".docker/"}

func BuildParentImage(path string) (int, error) {
    errCode := 0 
    var err error
   
    dockerImage, tag := searchInDockerfile(path)

    if dockerImage == "" {
        log.Printf("No docker image found.")
        return 1, errors.New("No Docker image found.")
    }

    err = os.Chdir("../" + defaultRepositoryName)

    if err != nil {
        err = os.Chdir("../")
        utils.GitClone(defaultRepository)    
        err = os.Chdir(defaultRepositoryName)
    } else {
        utils.GitPull()
    }
    utils.GitCheckout(dockerImage)
    Build(fmt.Sprintf("%s/%s:%s",defaultRegistryRepo, dockerImage, tag))

    return errCode, err
}

func searchInDockerfile(path string) (dockerImage, tag string) {
    paths := defaultPath
    if path != "" {
        paths = []string{path}
    }

    for _, path := range paths {
        fmt.Println("Search in " + path + dockerfile)
        if content, err := os.Open(path + dockerfile); err == nil {
            scanner := bufio.NewScanner(content)
            for scanner.Scan() {
                dockerImage, tag = extractImageName(scanner.Text())
                if dockerImage != "" {
                    fmt.Printf("Image found %s\n", dockerImage)
                    return dockerImage, tag
                }
            }
        }
    }

    return dockerImage, tag
}

/*
    Extract image name if is a ghcr.io/mylittleparis image
    From:
    FROM ghcr.io/mylittleparis/devops-docker-php/php7.4-fpm-alpine:latest AS prod
    To:
    php7.4-fpm-alpine:latest
*/
func extractImageName(line string) (dockerImage, tag string) {
    if dockerImage, found := strings.CutPrefix(line, "FROM " + defaultRegistry + "/" + defaultRepositoryName + "/"); found { 
        if found {fmt.Printf("Parent Image found %s\n", dockerImage)}
        if removeAs, _, found := strings.Cut(dockerImage, " AS "); found {
            if found {fmt.Printf("Image found %s\n", removeAs)}
            dockerImage, tag, _ = strings.Cut(removeAs, ":")
            return dockerImage, tag
        }
    }
    return dockerImage, tag
}

func Build(tag string)  {
    dockerBuild := exec.Command("docker", "build", "-t", tag, ".")
    fmt.Printf("Docker build %s", tag)
    pipe, _ := dockerBuild.StdoutPipe()

    err := dockerBuild.Start()

    if err != nil {
        log.Fatal(err)
    }

    reader := bufio.NewReader(pipe)
    line, err := reader.ReadString('\n')
    for err == nil {
        fmt.Println(line)
        line, err = reader.ReadString('\n')
    }

    err = dockerBuild.Wait()

    if err != nil {
        log.Printf("Command finished with error: %v", err)
    }
}
