package demo

import (
	"fmt"
	"sort"
	"time"
)

type TalkerDiscoveryPage struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	ModuleType int32     `json:"module_type"`
	Weight     int32     `json:"weight"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func main() {
	pages := []*TalkerDiscoveryPage{
		{ID: 1, Name: "Page 1", ModuleType: 1, Weight: 10},
		{ID: 2, Name: "Page 2", ModuleType: 2, Weight: 20},
		{ID: 3, Name: "Page 3", ModuleType: 3, Weight: 30},
		{ID: 4, Name: "Page 4", ModuleType: 1, Weight: 5},
		{ID: 5, Name: "Page 5", ModuleType: 2, Weight: 15},
		{ID: 6, Name: "Page 6", ModuleType: 4, Weight: 25},
		{ID: 7, Name: "Page 7", ModuleType: 1, Weight: 8},
		{ID: 8, Name: "Page 8", ModuleType: 4, Weight: 18},
		{ID: 9, Name: "Page 9", ModuleType: 3, Weight: 22},
	}

	sort.Slice(pages, func(i, j int) bool {
		// 先按 moduleType 的优先级排序
		priority := func(moduleType int32) int {
			switch moduleType {
			case 1:
				return 1 // 优先级最高
			case 2:
				return 2
			default:
				return 3 // 3 和 4 的优先级相同
			}
		}

		// 比较优先级
		pi, pj := priority(pages[i].ModuleType), priority(pages[j].ModuleType)
		if pi != pj {
			return pi < pj
		}

		// 同优先级按 weight 从大到小排序
		return pages[i].Weight > pages[j].Weight
	})

	// 打印排序后的结果
	for _, page := range pages {
		fmt.Printf("ID: %d, Name: %s, ModuleType: %d, Weight: %d\n", page.ID, page.Name, page.ModuleType, page.Weight)
	}
}
