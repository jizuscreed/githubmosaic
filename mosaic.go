package main

import (
	_ "fmt"
	"log"
	"fmt"
	_ "flag"
	"os"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"lib"
)

func getImagePath() (string, error) {
	//var path string
	if len(os.Args) > 1{
		return os.Args[1], nil
	} else {
		return "", errors.New("")
	}
}

/**
честно стырено
 */
func init() {
	// damn important or else At(), Bounds() functions will
	// caused memory pointer error!!
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

/**
дальшнейшие действия :
+ 1) вытащить работу с гитом в отдельный файл
2) там сделать отдельную структуру - gitrepo, в которой будет храниться путь к репе и прочее, чтобы не дрочиться постоянно. Сделать фукнции init repo и makeCommit
3) вытащить в отдельный файл разбор картинки
4) сделать определение дней по каледарю
5) итоговое создание репозитория с коммитами
 */

func main() {

	// получаем путь до картинки
	imagePath, err := getImagePath()

	if err != nil {
		log.Fatal("Please, specify image file")
	}

	// сначала проверяем, что мы можем работать с гитом
	lib.CheckGit()

	// пытемся открыть файл
	img := lib.OpenFile(imagePath)
	defer img.Close()

	fmt.Println(img)

	// так-с, запускаем анализ изображения
	calendarData := lib.GetCalendarDataFromImage(img)

	fmt.Println(calendarData.Result())

	// отлично, перезодим к созданию репозитория гита, мать его
	gitrepo := lib.NewGitRepo("test")

	fmt.Println(gitrepo)


	// defer // вот тут вот будет отлов исключений, ноя сука не помню щас, как это делается, так что потом, блять
}
