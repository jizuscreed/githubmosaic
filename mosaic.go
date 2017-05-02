package main

import (
	"os/exec"
	"bytes"
	_ "fmt"
	"strings"
	"log"
	"fmt"
	_ "flag"
	"os"
	"errors"
	"path/filepath"
	"image"
	"image/jpeg"
	"image/png"
)

const DAYS = 7
const WEEKS = 52
const MAX_COLOR = 65536; // uint16
const MAX_PERCENTAGE = MAX_COLOR * 3

type DayData struct {
	Sum int
	Counted int
	result int
	resultCounted bool
}

func (day *DayData) Add (red, green, blue uint32) {
	// собственно сначала считаем, какой процент буде мприбавлять
	newPercent := int(100 * (red + green + blue) / MAX_PERCENTAGE)
	// дальше неистово плюсуем процентаж и иинкрементим счетчик, ебать его в качель
	day.Sum += newPercent
	day.Counted++
}

func (day *DayData) Result() int {
	// так, нифига ещё не посчитано - ну значит считаем
	if !day.resultCounted {
		day.result = 100 - (day.Sum / day.Counted)
		day.resultCounted = true
	}
	return day.result
}

type WeekData struct {
	Days [DAYS]DayData
}

func (w *WeekData) Result() [DAYS]int {
	var result [DAYS]int
	for i := 0; i <len(w.Days); i++ {
		result[i] = w.Days[i].Result()
	}

	return result
}

type CalendarData struct {
	Weeks [WEEKS]WeekData
}

func (c *CalendarData) Result() [WEEKS][DAYS]int {
	var result [WEEKS][DAYS]int
	for i := 0; i <len(c.Weeks); i++ {
		result[i] = c.Weeks[i].Result()
	}

	return result
}

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

func getImagePath() (string, error) {
	//var path string
	if(len(os.Args) > 1){
		return os.Args[1], nil
	} else {
		return "", errors.New("")
	}
}

func openFile(filePath string) *os.File {
	// сначала вычисляем абсолютный путь к файлу
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if(err != nil){
		log.Fatal("Something went wrong");
	}
	// теперь пытаемся открыть картинку
	img, err := os.Open(dir + "/" + filePath)
	if(err != nil){
		log.Fatal(err.Error())
	}

	return img
}

func getCalendarDataFromImage(imgFile *os.File) CalendarData {
	calendarData := CalendarData{}

	// сначала получаем размер картинки
	imgCfg, _, err := image.DecodeConfig(imgFile)
	if err != nil{
		log.Fatal("Can not decode image file")
	}
	imgWidth, imgHeight := imgCfg.Width, imgCfg.Height
	// потрачено, начинаем читать из файла ()для начала перематываем назад ридер из файла, а то всё вальнется
	imgFile.Seek(0, 0)
	// получаем декодированное изображение
	img, _, _ := image.Decode(imgFile)

	// определяем максимальные точки перебора
	var maxWidth int = imgWidth - imgWidth % WEEKS
	var maxHeight int = imgHeight - imgHeight % DAYS


	var dayWidth int = int(maxWidth / WEEKS)
	var dayHeight int = int(maxHeight / DAYS)

	// перебираем пиксели и зажигаем пионерские костры
	for y := 0; y < maxHeight; y++ { // ползем по высоте
		for x := 0; x < maxWidth; x++ { // ползем по ширине
			r, g, b, _ := img.At(x, y).RGBA()
			// значицца тут индексами вычисляем, в какой день и какую неделю всё это класть и кладем
			calendarData.Weeks[int(x/dayWidth)].Days[int(y/dayHeight)].Add(r, g, b)
		}
	}

	return calendarData
}


func init() {
	// damn important or else At(), Bounds() functions will
	// caused memory pointer error!!
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

func main() {

	// получаем путь до картинки
	imagePath, err := getImagePath()

	if err != nil {
		log.Fatal("Please, specify image file")
	}

	// сначала проверяем, что мы можем работать с гитом
	checkGit();

	// пытемся открыть файл
	img := openFile(imagePath)
	defer img.Close()

	fmt.Println(img)

	// так-с, запускаем анализ изображения
	calendarData := getCalendarDataFromImage(img)

	fmt.Println(calendarData.Result())


	// defer // вот тут вот будет отлов исключений, ноя сука не помню щас, как это делается, так что потом, блять
}
