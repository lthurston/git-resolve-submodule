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
		fmt.Printf("usage: git resolve-submodule <pathish>\n")
		os.Exit(1)
	}

	// First, confirm that we're within a git repo
	mustBeWithinRepo()

	switch os.Args[1] {
	case "..":
		cmDotDot()
	case ".":
		cmDot()
	case "/":
		cmSlash()
	default:
		cmFind(os.Args[1])
	}

	path, _ := os.Getwd()
	fmt.Print(path)
}

func mustBeWithinRepo() {
	err := withinRepo()
	if err != nil {
		log.Fatal(err)
	}
}

func withinRepo() (err error) {
	_, err = exec.Command("git", "status", "--porcelain").Output()
	return
}

func cmDotDot() {
	cmDot()
	os.Chdir("..")
	cmDot()
}

func cmDot() {
	err := chdirToCurrentRepoRoot()
	if err != nil {
		log.Fatal(err)
	}
}

func chdirToCurrentRepoRoot() error {
	path, err := getToplevel()
	if err != nil {
		return err
	}
	pathString := strings.TrimSpace(string(path))
	err = os.Chdir(pathString)
	return err
}

func cmSlash() {
	var dir string
	err := chdirToCurrentRepoRoot()
	if err != nil {
		log.Fatal(err)
	}
	dir, err = os.Getwd()
	for err == nil {
		dir, _ = os.Getwd()
		os.Chdir("..")
		chdirToCurrentRepoRoot()
		_, err = getToplevel()
	}

	os.Chdir(dir)
}

func getToplevel() ([]byte, error) {
	return exec.Command("git", "rev-parse", "--show-toplevel").Output()
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
