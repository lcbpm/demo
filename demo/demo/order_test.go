package demo

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

type DiscoveryPageModuleType int32

const (
	DiscoveryPageModuleType_DPMT_None       DiscoveryPageModuleType = 0
	DiscoveryPageModuleType_DPMT_Banner     DiscoveryPageModuleType = 1
	DiscoveryPageModuleType_DPMT_Functional DiscoveryPageModuleType = 2
	DiscoveryPageModuleType_DPMT_Content    DiscoveryPageModuleType = 3
	DiscoveryPageModuleType_DPMT_Feature    DiscoveryPageModuleType = 4
)

type TalkerDiscoveryPage struct {
	ID         int64                   `json:"id"`
	Name       string                  `json:"name"`
	ModuleType DiscoveryPageModuleType `json:"module_type"`
	Weight     int32                   `json:"weight"`
	CreatedAt  time.Time               `json:"created_at"`
	UpdatedAt  time.Time               `json:"updated_at"`
}

func Test_rder_test(t *testing.T) {
	pages := []*TalkerDiscoveryPage{
		{ID: 1, Name: "Page 1", ModuleType: DiscoveryPageModuleType_DPMT_Banner, Weight: 10},
		{ID: 2, Name: "Page 2", ModuleType: DiscoveryPageModuleType_DPMT_Functional, Weight: 20},
		{ID: 3, Name: "Page 3", ModuleType: DiscoveryPageModuleType_DPMT_Content, Weight: 30},
		{ID: 4, Name: "Page 4", ModuleType: DiscoveryPageModuleType_DPMT_Banner, Weight: 5},
		{ID: 5, Name: "Page 5", ModuleType: DiscoveryPageModuleType_DPMT_Functional, Weight: 15},
		{ID: 6, Name: "Page 6", ModuleType: DiscoveryPageModuleType_DPMT_Feature, Weight: 25},
		{ID: 7, Name: "Page 7", ModuleType: DiscoveryPageModuleType_DPMT_Banner, Weight: 8},
		{ID: 8, Name: "Page 8", ModuleType: DiscoveryPageModuleType_DPMT_Feature, Weight: 18},
		{ID: 9, Name: "Page 9", ModuleType: DiscoveryPageModuleType_DPMT_Content, Weight: 22},
	}

	sort.Slice(pages, func(i, j int) bool {
		// 先按 moduleType 的优先级排序
		priority := func(moduleType DiscoveryPageModuleType) int {
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
