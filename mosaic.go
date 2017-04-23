package main

import (
	"os/exec"
	"bytes"
	_ "fmt"
	"strings"
	"log"
)

func checkGit() {
	// проверяем, что гит есть в системе
	cmd := exec.Command("git", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	// проверяем ответ
	if(!strings.Contains(out.String(), "git version")){
		log.Fatal("Can not find git. Please install and try again")
	}
}

func main() {
	//imgFile := "1.jpg" // todo потом переделать на получение этого дельца из консоли

	// сначала проверяем, что мы можем работать с гитом
	checkGit();
	// пытемся открыть файл



	// defer // вот тут вот будет отлов исключений, ноя сука не помню щас, как это делается, так что потом, блять
}
