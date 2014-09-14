package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: git resolve-submodule <module pathish>\n")
		os.Exit(1)
	}

	// First, confirm that we're within a git repo
	mustBeWithinRepo()

	switch os.Args[1] {
	case "..":
		cmDotDot()
	case ".":
		cmDot()
	default:
		cmFind(os.Args[1])
	}

	path, _ := os.Getwd()
	fmt.Print(path)
}

func mustBeWithinRepo() {
	_, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		log.Fatal(err)
	}
}

func cmDotDot() {
	cmDot()
	os.Chdir("..")
	cmDot()
	mustBeWithinRepo()
}

func cmDot() {
	path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		log.Fatal(err)
	}
	pathString := strings.TrimSpace(string(path))
	err = os.Chdir(pathString)
}

func cmFind(find string) {
	path, err := exec.Command("git", "submodule").Output()
	if err != nil {
		log.Fatal(err)
	}

	pathString := strings.TrimSpace(string(path))
	pathLines := strings.Split(pathString, "\n")

	find = strings.ToUpper(find)
	for _, line := range pathLines {
		line = strings.TrimSpace(line)
		lineFields := strings.Fields(line)
		if strings.Contains(strings.ToUpper(line), find) == true {
			err = os.Chdir(lineFields[1])
			return
		}
	}
}
