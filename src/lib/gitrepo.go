package lib

import (
	"os"
	"log"
	_ "fmt"
	"errors"
	"os/exec"
)

/**
 дескриптор репозитория гита
 */
type gitrepo struct {
	dir               string
	initialized       bool
	unCommitedChanges bool
	lastCommitContent string
}

func (repo gitrepo) GetDir() string {
	return repo.dir
}

func (repo gitrepo) IsInitialized() bool {
	return repo.initialized
}

func (repo gitrepo) HasUnCommitedChanges() bool {
	return repo.unCommitedChanges
}

func (repo *gitrepo) Init() error {
	// так-с, проверяем, что репа ещё не инициализирована
	if repo.initialized{
		return errors.New("git repository in " + repo.dir + "is already initinalized")
	}

	// отлично, инициализируемся
	exec.Command("git", "init", repo.dir).Run()

	// теперь создаём файл
	file, err := os.Create(repo.dir + "/output.txt")
	// закрываем файл
	defer file.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

func (repo *gitrepo) NewCommit(date string) error {
	// меняем содержимое файла
	file, err := os.OpenFile("output.txt", os.O_TRUNC|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err.Error())
	}
	// пишем туда
	if repo.lastCommitContent == "+" {
		_, err = file.Write([]byte("-"))
	} else {
		_, err = file.Write([]byte("+"))
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	// закрываем файл
	file.Close()

	// добавляем изменения в индекс (git add)
	exec.Command("git", "add", repo.dir).Run()

	// коммитимся (git commit -q --date "2017-01-01 12:00:02" -m "+")
	exec.Command("git", "commit", repo.dir, "-q", "--date", "\"" + date + "\"", "-m", "\"+\"").Run()

	return nil
}

func NewGitRepo(repoDir string) gitrepo {
	currentDir, _ := getCurrentDir()
	abspath := currentDir + "/../" + repoDir

	// так-с, пиздуем в нужную директорию, если её не существует, то создаём
	_, err := os.Stat(abspath)

	if os.IsNotExist(err) {
		// отлично, директории не существует - создаём её и работаем дальше
		os.Mkdir(abspath, 0777)
	} else if err != nil{
		log.Fatal(err.Error())
	} else {
		// директория уже существует - бросеам фаталку
		//log.Fatal("git repo directory " + repoDir + " is already exists") todo не забыть раскомметировать потом, мазафака
	}

	repo := gitrepo{abspath, false, false, "-"}
	repo.Init()

	return repo
}