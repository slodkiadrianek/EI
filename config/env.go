package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Env struct {
	Port   string
	DbLink string
}

func readFile(pathToFile string) map[string]string {
	file, err := os.OpenFile(pathToFile, os.O_RDONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	envVariables := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplitted := strings.SplitN(line, "=", 2)
		fmt.Println(lineSplitted)
		envVariables[lineSplitted[0]] = lineSplitted[1]
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return envVariables
}

func SetConfig(pathToFile string) *Env {
	envVariables := readFile(pathToFile)
	return &Env{
		Port:   envVariables["PORT"],
		DbLink: envVariables["DbLink"],
	}
}
