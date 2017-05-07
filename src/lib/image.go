package lib

import (
	"path/filepath"
	"os"
	"log"
	"image"
)
/**
 возвращает текущую директорию
 */
func getCurrentDir() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil{
		return "", err
	}

	return dir, nil
}

/**
 открывает изображение для чтения
 */
func OpenFile(filePath string) *os.File {
	// сначала вычисляем абсолютный путь к файлу
	dir, err := getCurrentDir()
	if err != nil{
		log.Fatal(err.Error())
	}
	// теперь пытаемся открыть картинку
	img, err := os.Open(dir + "/" + filePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	return img
}
/**
 получает каледарные данные из анализа изображения
 */
func GetCalendarDataFromImage(imgFile *os.File) CalendarData {
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
