package main
import (
	"fmt"
	"strconv"
	"strings"
	"sort"
  "math"
)

func Stati(strg string) string {
	arrayOfResults := strings.Split(strg, ", ")
	arrayOfSeconds := make([]int, 0)
  if len(arrayOfResults) == 1 {
    return ""
  }
	for _, result := range arrayOfResults {
		resultSplit    := strings.Split(result, "|")
		hour, _        := strconv.Atoi(resultSplit[0])
		minute, _      := strconv.Atoi(resultSplit[1])
		second, _      := strconv.Atoi(resultSplit[2])
		arrayOfSeconds = append(arrayOfSeconds, CountSeconds(hour, minute, second))
	}
	for _, elem := range arrayOfSeconds {
		fmt.Println(HourMinSec(elem))
	}
	return  Range(arrayOfSeconds) + " " + MeanOrAverage(arrayOfSeconds) + " " + Median(arrayOfSeconds)
}

func CountSeconds(h, m, s int) int {
	return h*3600 + m*60 + s
}

func Range(arrayOfSeconds []int) string {
	sort.Ints(arrayOfSeconds)
	min := arrayOfSeconds[0]
	max := arrayOfSeconds[len(arrayOfSeconds) - 1]
	rangeValue := max-min
	return "Range: " +  HourMinSec(rangeValue)
}

func MeanOrAverage(arrayOfSeconds []int) string {
	sum := 0
	for _, elem := range arrayOfSeconds {
		sum += elem
	}
	return "Average: " + HourMinSec(sum/len(arrayOfSeconds))
}

func Median(arrayOfSeconds []int) string {
	var resultValue int
	var length int = len(arrayOfSeconds)
	if len(arrayOfSeconds) % 2 == 0 {
		resultValue = (arrayOfSeconds[length/2-1] + arrayOfSeconds[length/2])/2
	} else {
		resultValue = arrayOfSeconds[length/2]
	} 
	return "Median: " + HourMinSec(resultValue)
}

func AddZero(str string) string {
	if len(str) == 1 {
		return "0" + str
	} else {
		return str
	}
}

func HourMinSec(secs int) string {
	hours    := float64(secs) / 3600
	minutes  := (float64(secs)/3600 - math.Floor(hours))*60
	seconds  := math.Round((minutes - math.Floor(minutes))*60)
	if seconds == 60.0 {
		minutes += 1
		seconds  = 0
	}
	if minutes == 60.0 {
		hours  += 1
		minutes = 0
	}
	return AddZero(strconv.Itoa(int(hours))) + "|" + AddZero(strconv.Itoa(int(math.Floor(minutes)))) + "|" + AddZero(strconv.Itoa(int(seconds)))
}

func main() {
	fmt.Println(Stati("02|15|59, 2|47|16, 02|17|20, 2|32|34, 2|17|17, 2|22|00, 2|31|41"))
}
