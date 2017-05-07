package lib

import (
	"os"
	"log"
	"fmt"
)

/**
 дескриптор репозитория гита
 */
type gitrepo struct {
	dir string
	initialized bool
	unCommitedChanges bool
}

func (repo gitrepo) GetDir() string {
	return repo.dir
}

func (repo gitrepo) IsInitialized() bool {
	return repo.initialized;
}

func (repo gitrepo) HasUnCommitedChanges() bool {
	return repo.unCommitedChanges
}

func NewGitRepo(repoDir string) gitrepo {
	currentDir, _ := getCurrentDir()

	// так-с, пиздуем в нужную директорию, если её не существует, то создаём
	fileinfo, err := os.Stat(currentDir + "/../" + repoDir);

	if err != nil{
		log.Fatal(err.Error())
	}

	fmt.Println(fileinfo)

	repo := gitrepo{repoDir, false, false}

	return repo
}