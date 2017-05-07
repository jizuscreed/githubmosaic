package lib

import (
	"os/exec"
	"bytes"
	"strings"
	"log"
)

/**
проверяет доступность гита
 */
func CheckGit() {
	// проверяем, что гит есть в системе
	cmd := exec.Command("git", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	// проверяем ответ
	if !strings.Contains(out.String(), "git version") {
		log.Fatal("Can not find git. Please install and try again")
	}
}