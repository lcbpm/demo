package time

import (
	"fmt"
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
}
