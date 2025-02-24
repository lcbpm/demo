package lang

import (
	"fmt"
	"github.com/longbridgeapp/opencc" // 引入 OpenCC 库，需提前安装
	"testing"
	"unicode"
)

// DetectLanguageResult 用于存储检测结果
type DetectLanguageResult struct {
	Language   string
	Script     string
	Confidence float64
}

// DetectLanguage 检测输入文本的语言和文字类型（简体/繁体）
func DetectLanguage(text string) (DetectLanguageResult, error) {
	var totalLetters, hanLetters, latinLetters int

	// 遍历文本，统计字母字符数量
	for _, r := range text {
		if unicode.IsLetter(r) {
			totalLetters++
			if unicode.In(r, unicode.Han) {
				hanLetters++
			} else if unicode.In(r, unicode.Latin) {
				latinLetters++
			}
		}
	}

	// 避免除以零
	if totalLetters == 0 {
		return DetectLanguageResult{
			Language:   "未知",
			Script:     "未知",
			Confidence: 0.0,
		}, nil
	}

	// 计算汉字和拉丁字母的比例
	hanRatio := float64(hanLetters) / float64(totalLetters)
	latinRatio := float64(latinLetters) / float64(totalLetters)

	result := DetectLanguageResult{}

	if hanRatio > 0.5 {
		// 主要为中文

		// 初始化 OpenCC 转换器
		converterS2T, err := opencc.New("s2t.json")
		if err != nil {
			return result, fmt.Errorf("初始化简转繁转换器失败：%v", err)
		}
		converterT2S, err := opencc.New("t2s.json")
		if err != nil {
			return result, fmt.Errorf("初始化繁转简转换器失败：%v", err)
		}

		// 对整个文本进行简繁转换
		toTraditional, err := converterS2T.Convert(text)
		if err != nil {
			return result, fmt.Errorf("简转繁转换失败：%v", err)
		}
		toSimplified, err := converterT2S.Convert(text)
		if err != nil {
			return result, fmt.Errorf("繁转简转换失败：%v", err)
		}

		// 判断简体或繁体
		if text == toSimplified && text != toTraditional {
			result.Script = "简体中文"
		} else if text == toTraditional && text != toSimplified {
			result.Script = "繁体中文"
		} else if text == toSimplified && text == toTraditional {
			result.Script = "简体和繁体相同"
		} else {
			result.Script = "简繁混合或无法区分"
		}

		result.Language = "中文"
		result.Confidence = hanRatio

	} else if latinRatio > 0.5 {
		// 主要为英文
		result.Language = "英文"
		result.Script = "拉丁字母"
		result.Confidence = latinRatio

	} else {
		// 其他语言或无法识别
		result.Language = "未知"
		result.Script = "未知"
		result.Confidence = 0.0
	}

	return result, nil
}

func Test_open(t *testing.T) {
	// 示例文本，可以替换为您的输入
	texts := []string{
		"你好，世界！",        // 简体中文
		"妳好，世界！",        // 繁体中文
		"您好，世界！",        // 简繁相同
		"Hello, world!", // 英文
		"你好，world！",     // 中英文混合
		"",              // 空字符串
		"12345",         // 纯数字
		"!@#$%^&*()",    // 符号
	}

	for _, text := range texts {
		result, err := DetectLanguage(text)
		if err != nil {
			fmt.Printf("检测文本 \"%s\" 时出错：%v\n", text, err)
			continue
		}
		fmt.Printf("文本：%s\n", text)
		fmt.Printf("语言：%s\n", result.Language)
		fmt.Printf("文字类型：%s\n", result.Script)
		fmt.Printf("置信度：%.2f%%\n\n", result.Confidence*100)
	}
}
