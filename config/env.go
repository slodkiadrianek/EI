package config

import (
	"bufio"
	"os"
	"strings"
)

type Env struct {
	Port         string
	DbLink       string
}

func readFile() map[string]string {
	file, err := os.OpenFile(".env", os.O_RDONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	envVariables := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplitted := strings.SplitN(line, "=", 2)
		envVariables[lineSplitted[0]] = lineSplitted[1]
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return envVariables
}

func SetConfig() *Env {
	envVariables := readFile()
	return &Env{
		Port:         envVariables["Port"],
		DbLink:       envVariables["DbLink"],
	}
}