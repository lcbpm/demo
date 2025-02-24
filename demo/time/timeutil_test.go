package time

import (
	"fmt"
	"github.com/gookit/goutil/timex"
	"testing"
	"time"
)

func daysBetweenTimestamps(startTimestamp, endTimestamp int64) int {
	// 将时间戳转换为 time.Time
	start := time.Unix(startTimestamp, 0)
	end := time.Unix(endTimestamp, 0)

	// 计算日期范围内的天数
	duration := end.Sub(start)
	days := int(duration.Hours() / 24)
	return days
}

func printDatesInRange(startTimestamp, endTimestamp int64) []string {
	var result []string
	start := time.Unix(startTimestamp, 0)
	end := time.Unix(endTimestamp, 0)

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		result = append(result, d.Format("20060102"))
	}
	return result
}

func getDateByRegion() {
	loc := GetUTCLocationByRegion("CN")
	GetZeroTimeInLoc(time.Now().In(loc), loc).Unix()
	loc = GetUTCLocationByRegion("CN")
	GetZeroTimeInLoc(time.Now().In(loc), loc).Unix()
	time.Now().Unix()
}

func getDateByRegion2() {
	loc := GetUTCLocationByRegion("cn")
	layout := "15:04"
	parsedTime, err := time.ParseInLocation(layout, "11:04", loc)
	if err != nil {
	}
	fmt.Printf("%v\n", parsedTime.Minute())
}

func Test_time(t *testing.T) {
	fmt.Printf("%v\n", time.Now().Unix())
	fmt.Printf("%v", time.Now().AddDate(0, 0, 7).Unix())

	fmt.Printf("%v\n", daysBetweenTimestamps(time.Now().Unix(), time.Now().AddDate(0, 0, 7).Unix()))
	fmt.Printf("%v\n", printDatesInRange(time.Now().Unix(), time.Now().AddDate(0, 0, 7).Unix()))
	getDateByRegion2()
	fmt.Printf("%v\n", time.Now().Truncate(time.Minute))

	tx, _ := timex.FromDate("2024-02-10 00:00:00")
	currentTime := tx.Time
	//nextMonthSameTime := common.GetNextMonthSameDaySameTime(currentTime)
	sameTime, i := GetFutureMonthSameDaySameTime(currentTime, 1)
	daySameTime, i2 := GetNextMonthSameDaySameTime(currentTime)
	fmt.Printf("%v,%v\n", sameTime, i)
	fmt.Printf("%v,%v", daySameTime, i2)

}

func GetNextMonthSameDaySameTime(currentTime time.Time) (time.Time, int64) {
	// 当前时间
	currentYear, currentMonth, currentDay := currentTime.Date()
	currentHour, currentMinute, currentSecond := currentTime.Clock()

	// 下个月的时间
	nextMonth := currentMonth + 1
	nextYear := currentYear

	// 考虑是否需要跨年和跨月处理
	if nextMonth > 12 {
		nextMonth = 1
		nextYear++
	}

	// 获取下个月的第一天
	nextMonthFirstDay := time.Date(nextYear, time.Month(nextMonth), 1, currentHour, currentMinute, currentSecond, 0, currentTime.Location())

	// 获取当前日期在下个月的对应日期
	nextMonthSameDay := currentDay
	if nextMonthSameDay > 28 {
		// 如果当前日期大于28号，可能需要调整到合法日期
		nextMonthLastDay := nextMonthFirstDay.AddDate(0, 1, -1).Day()
		if nextMonthSameDay > nextMonthLastDay {
			nextMonthSameDay = nextMonthLastDay
		}
	}

	// 组合成完整时间
	nextMonthSameTime := time.Date(nextYear, time.Month(nextMonth), nextMonthSameDay, currentHour, currentMinute, currentSecond, 0, currentTime.Location())
	return nextMonthSameTime, int64(nextMonthSameTime.Sub(currentTime).Hours() / 24)
}

func GetFutureMonthSameDaySameTime(currentTime time.Time, monthsAhead int) (time.Time, int64) {
	// 当前时间
	currentYear, currentMonth, currentDay := currentTime.Date()
	currentHour, currentMinute, currentSecond := currentTime.Clock()

	// 计算目标的年和月
	totalMonths := int(currentMonth) + monthsAhead
	futureYear := currentYear + (totalMonths-1)/12
	futureMonth := (totalMonths-1)%12 + 1

	// 获取目标月份的第一天
	futureMonthFirstDay := time.Date(futureYear, time.Month(futureMonth), 1, currentHour, currentMinute, currentSecond, 0, currentTime.Location())

	// 获取当前日期在目标月份的对应日期
	futureDay := currentDay
	if futureDay > 28 {
		// 如果当前日期大于28号，可能需要调整到合法日期
		futureMonthLastDay := futureMonthFirstDay.AddDate(0, 1, -1).Day()
		if futureDay > futureMonthLastDay {
			futureDay = futureMonthLastDay
		}
	}

	// 组合成完整时间
	futureSameTime := time.Date(futureYear, time.Month(futureMonth), futureDay, currentHour, currentMinute, currentSecond, 0, currentTime.Location())
	return futureSameTime, int64(futureSameTime.Sub(currentTime).Hours() / 24)
}
