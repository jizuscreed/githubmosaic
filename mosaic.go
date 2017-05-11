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
	"time"
	"strconv"
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
1) сделать всю работу с календарем (там оказывается совсем не сложно)
+ 2) дописать создание коммитов
+ 2,1) работа с файлом - писанина туда плючиков/минусиков, ёпта (с этого вообще стоило бы начать, а то это и для коммита надои для всего остального прочего)
3) тестирование
4) profit
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

	calendarResult := calendarData.Result()

	// отлично, переходим к созданию репозитория гита, мать его
	gitrepo := lib.NewGitRepo("test")

	// дальше получем дату на год назад и и топаем вперед до понедельника
	date := time.Now()
	date = date.AddDate(-1, 0, 0)
	for date.Weekday().String() != "Monday" {
		date = date.AddDate(0, 0, 1)
	}

	// теперь округляем дату - отбрасываем нафиг часы, минуты и секунды
	date = date.Add(time.Duration(-1*date.Second()) * time.Second)
	date = date.Add(time.Duration(-1*date.Minute()) * time.Minute)
	date = date.Add(time.Duration(-1*date.Hour()) * time.Hour)

	// отлично, мы нашли стартовый день - теперь перебираем данные из календарных данных и фигачим коммиты
	for i := 0; i < len(calendarResult); i++{ // недели
		fmt.Println("week " + strconv.Itoa(i))
		for l := 0; l < len(calendarResult[i]); l++ { // дни в неделе
			fmt.Println("day " + strconv.Itoa(l))
			tempDate := date // создаём временную дату, чтобы там изгаляться с секундами
			// теперь делаем столько коммитов, сколько у нас процент заполненности дня
			// (это для простоты, чтобы один процентаж не перегонять в другой)
			for m := 0; m <calendarResult[i][l]; m++{
				fmt.Println("commit " + strconv.Itoa(m))
				gitrepo.NewCommit(tempDate.Format("2006-01-02 15:04:05"))
				tempDate = tempDate.Add(time.Second) // переводим временную дату на секунду вперед
			}
			// отлично, коммиты за день готовы - топаем на день вперед
			date = date.AddDate(0, 0, 1)
		}
	}

	// по идее всё, можно тестировать
}
