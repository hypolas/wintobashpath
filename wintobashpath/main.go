package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//go:embed bashrc
var bashrc []byte

func main() {
	var winPathList []string
	path := os.Getenv("PATH")
	splitedPath := strings.Split(path, ";")
	for _, p := range splitedPath {
		pathSlach := filepath.ToSlash(p)

		winPathList = append(winPathList, fmt.Sprintf("%q", renameDrive(pathSlach)))
	}

	bashrcPath := strings.Join(winPathList, ":")
	userProfile := os.Getenv("USERPROFILE")

	if !exists(userProfile + "/.bashrc.d") {
		createBashrcDir(userProfile + "/.bashrc.d")
	}

	os.WriteFile(userProfile+"/.bashrc.d/winToBashPath.bashrc", []byte("export PATH=$PATH:"+bashrcPath), 0700)

	if !exists(userProfile + "/.bashrc") {
		os.WriteFile(userProfile+"/.bashrc", bashrc, 0700)
	} else {
		readBashRc, _ := os.ReadFile(userProfile + "/.bashrc")
		if !strings.Contains(string(readBashRc), "winToBashPath") {
			f, err := os.OpenFile(userProfile+"/.bashrc",
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0700)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			f.WriteString(string(bashrc))
			f.Close()
		}
	}
}

func spaceEscaped(s string) string {
	return strings.Replace(s, " ", "\\ ", -1)
}

func renameDrive(path string) string {
	var drive string = strings.Split(path, ":")[0]
	path = strings.Replace(path, drive+":", "/"+strings.ToLower(drive), 1)
	return path
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func createBashrcDir(path string) {
	if err := os.Mkdir(path, 0700); err != nil {
		log.Fatal(err)
	}
}
