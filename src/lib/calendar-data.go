package lib

import "math"

const DAYS = 7
const WEEKS = 52
const MAX_COLOR = 65536; // uint16
const MAX_PERCENTAGE = MAX_COLOR * 3

/**
 данные одного дня
 */
type DayData struct {
	Sum int
	Counted int
	result int
	resultCounted bool
}

func (day *DayData) Add (red, green, blue uint32) {
	// собственно сначала считаем, какой процент будем прибавлять (округляем вверх, а не вниз)
	newPercent := int(math.Ceil(float64(100 * (red + green + blue)) / MAX_PERCENTAGE))
	// дальше неистово плюсуем процентаж и иинкрементим счетчик, ебать его в качель
	day.Sum += newPercent
	day.Counted++
}

func (day *DayData) Result() int {
	// так, нифига ещё не посчитано - ну значит считаем
	if !day.resultCounted {
		// сразу инвертируем
		day.result = 100 - (day.Sum / day.Counted)
		day.resultCounted = true
	}
	return day.result
}

/**
 данные недели
 */
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

/**
 данные за год
 */
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
