package time

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

// ConvertBeijingTimeToCustomTime 根据输入的时区名称转换北京时间到指定时区
func ConvertBeijingTimeToCustomTime(beijingTime time.Time, timezoneName string) (time.Time, error) {
	// 去除空格
	timezoneName = strings.TrimSpace(timezoneName)

	// 先尝试解析为 UTC±X 或 UTC±X:XX 格式
	offsetHours, offsetMinutes, err := ParseUTCOffset(timezoneName)
	if err == nil {
		// 解析成功，转换为目标时区
		return ConvertBeijingTimeWithOffset(beijingTime, offsetHours, offsetMinutes), nil
	}

	// 如果不是 UTC 格式，尝试解析为 IANA 时区
	location, err := time.LoadLocation(timezoneName)
	if err == nil {
		// 解析成功，返回转换后的时间
		return beijingTime.In(location), nil
	}

	// 如果是常见地方时区标识（如 CST, PST 等），则进行映射
	offsetHours, offsetMinutes, err = GetTimezoneOffsetFromAbbreviation(timezoneName)
	if err == nil {
		// 地方时区标识转换为对应的偏移
		return ConvertBeijingTimeWithOffset(beijingTime, offsetHours, offsetMinutes), nil
	}

	// 如果都无法解析，返回错误
	return time.Time{}, fmt.Errorf("无法识别的时区名称: %s", timezoneName)
}

// ConvertBeijingTimeWithOffset 将北京时间转换为给定的时区偏移
func ConvertBeijingTimeWithOffset(beijingTime time.Time, offsetHours, offsetMinutes int) time.Time {
	// 计算总的秒数偏移
	offset := offsetHours*3600 + offsetMinutes*60

	// 创建一个自定义时区
	location := time.FixedZone(fmt.Sprintf("UTC%+d:%02d", offsetHours, offsetMinutes), offset)

	// 将北京时间转换为目标时区的时间
	return beijingTime.In(location)
}

// ParseUTCOffset 解析 UTC±X 或 UTC±X:XX 格式的时区
func ParseUTCOffset(timezoneName string) (int, int, error) {
	// 正则解析 UTC±X 或 UTC±X:XX 格式
	re := regexp.MustCompile(`^UTC([+-])(\d{1,2})(?::(\d{2}))?$`)
	matches := re.FindStringSubmatch(timezoneName)

	if len(matches) < 3 {
		return 0, 0, fmt.Errorf("无效的 UTC 偏移格式: %s", timezoneName)
	}

	sign := matches[1]
	hours, err := strconv.Atoi(matches[2])
	if err != nil {
		return 0, 0, fmt.Errorf("无效的小时数: %s", matches[2])
	}

	minutes := 0
	if len(matches) > 3 && matches[3] != "" {
		minutes, err = strconv.Atoi(matches[3])
		if err != nil {
			return 0, 0, fmt.Errorf("无效的分钟数: %s", matches[3])
		}
	}

	// 根据正负号调整小时数
	if sign == "-" {
		hours = -hours
	}

	return hours, minutes, nil
}

// GetTimezoneOffsetFromAbbreviation 根据常见地方时区标识（如 CST, PST）返回对应的UTC偏移
func GetTimezoneOffsetFromAbbreviation(timezoneName string) (int, int, error) {
	// 常见地方时区标识和其对应的 UTC 偏移
	abbreviationMap := map[string]struct {
		hours   int
		minutes int
	}{
		"CST": {8, 0},  // China Standard Time (北京时间)
		"PST": {-8, 0}, // Pacific Standard Time
		"EST": {-5, 0}, // Eastern Standard Time
		"MST": {-7, 0}, // Mountain Standard Time
		"GMT": {0, 0},  // Greenwich Mean Time (UTC)
		"BST": {1, 0},  // British Summer Time
		"JST": {9, 0},  // Japan Standard Time
		"IST": {5, 30}, // Indian Standard Time
	}

	// 查找并返回对应的 UTC 偏移
	if offset, exists := abbreviationMap[timezoneName]; exists {
		return offset.hours, offset.minutes, nil
	}

	return 0, 0, fmt.Errorf("无法识别的地方时区标识: %s", timezoneName)
}

// 获取北京时间（CST）时间
func GetBeijingTime() time.Time {
	return time.Date(2024, 11, 12, 0, 21, 0, 0, time.FixedZone("CST", 8*3600))
}

func Test_zone(t *testing.T) {
	// 获取北京时间
	beijingTime := GetBeijingTime()

	// 示例: 转换为不同的时区名称
	timezoneNames := []string{
		"UTC-3",         // UTC-3
		"UTC-4:30",      // UTC-4:30
		"UTC+5:00",      // UTC+5:00
		"UTC+8",         // UTC+8 (北京时间)
		"Asia/Shanghai", // IANA时区
		"Europe/London", // IANA时区
		"PST",           // 地方时区
		"CST",           // 地方时区
		"GMT",           // GMT时区
		"IST",           // Indian Standard Time
		"JST",           // Japan Standard Time
	}

	// 遍历所有时区名称并输出转换后的时间
	for _, timezoneName := range timezoneNames {
		targetTime, err := ConvertBeijingTimeToCustomTime(beijingTime, timezoneName)
		if err != nil {
			fmt.Println("错误:", err)
			continue
		}
		fmt.Printf("北京时间: %s 转换为 %s 时间: %s\n", beijingTime.Format("2006-01-02 15:04:05"), timezoneName, targetTime.Format("2006-01-02 15:04:05"))
	}
}
