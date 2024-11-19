package main

import (
	"fmt"
	"testing"
	"time"
)

// 计算跳过周末的日期（包括当前天）
func skipWeekends(start time.Time, days int) time.Time {
	result := start
	for days > 0 {
		if result.Weekday() != time.Saturday && result.Weekday() != time.Sunday {
			days-- // 如果不是周末，计入目标天数
		}
		// 如果 days > 0 才继续加一天，避免多加一天
		if days > 0 {
			result = result.AddDate(0, 0, 1) // 每次加一天
		}
	}
	return result
}

// 计算包括周末的日期（包括当前天）
func includeWeekends(start time.Time, days int) time.Time {
	return start.AddDate(0, 0, days-1) // 减 1 因为当前天算作第 1 天
}

func Test_(t *testing.T) {
	// 获取当前时间
	location, err := time.LoadLocation("Asia/Shanghai") // 设置时区
	if err != nil {
		t.Fatalf("加载时区失败: %v", err)
	}

	start := time.Now().In(location) // 当前时间加上指定时区
	days := 7                        // 往后 7 天

	// 计算跳过周末的最终日期
	resultSkip := skipWeekends(start, days)
	fmt.Printf("跳过周末后的日期: %s\n", resultSkip.Format("2006-01-02 15:04:05"))

	// 计算包括周末的最终日期
	resultInclude := includeWeekends(start, days)
	fmt.Printf("包括周末后的日期: %s\n", resultInclude.Format("2006-01-02 15:04:05"))
}
