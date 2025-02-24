package lang

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/abadojack/whatlanggo"
)

// 检测文本语言（不依赖硬编码字符匹配）
func detectLanguage(text string) string {
	// 处理空输入或无效输入
	text = strings.TrimSpace(text)
	if text == "" {
		return "空输入"
	}
	if !containsLetters(text) {
		return "无有效字母"
	}

	// 使用 whatlanggo 初步检测
	info := whatlanggo.Detect(text)

	// 如果是中文
	if info.Lang == whatlanggo.Cmn {
		// 通过统计字符复杂度来推测简繁（简体字符通常更简单）
		return detectChineseType(text)
	}

	// 如果是英文
	if info.Lang == whatlanggo.Eng && info.Confidence > 0.7 {
		return "英文"
	}

	// 其他语言
	return fmt.Sprintf("其他语言 (%s, 置信度: %.2f)", info.Lang.String(), info.Confidence)
}

// 判断中文类型（简体或繁体）
func detectChineseType(text string) string {
	// 计算汉字的笔画数（近似方法：基于Unicode编码的统计特性）
	var hanCount, complexCount int
	for _, r := range text {
		if unicode.Is(unicode.Han, r) {
			hanCount++
			// 假设繁体字符更复杂，编码值更高或笔画数更多
			// 这是一种简化统计方法，实际可以用笔画数据库
			if isComplexChar(r) {
				complexCount++
			}
		}
	}

	if hanCount == 0 {
		return "中文（无汉字）"
	}

	// 如果复杂字符占比超过50%，认为是繁体
	complexRatio := float64(complexCount) / float64(hanCount)
	if complexRatio > 0.5 {
		return "繁体中文"
	} else if complexRatio > 0.1 {
		return "混合中文（简繁共存）"
	}
	return "简体中文"
}

// 判断字符是否“复杂”（简化为Unicode值或笔画数的估计）
func isComplexChar(r rune) bool {
	// 简单假设：繁体字符的Unicode值通常较高，或结构更复杂
	// 注意：这只是一个近似方法，实际需要笔画数据
	return r > 0x4E00 && (r > 0x9FA5 || (r >= 0x3400 && r <= 0x4DBF))
}

// 判断是否包含字母
func containsLetters(text string) bool {
	for _, r := range text {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func Test_opencc(t *testing.T) {
	// 测试用例
	tests := []string{
		"Hello, how are you?", // 纯英文
		"你好，今天怎么样？",           // 纯简体
		"妳好，今天怎麼樣？",           // 纯繁体
		"你好，今天怎麼樣？",           // 简繁混合
		"こんにちは",               // 日文
		"",                    // 空输入
		"123 !@#",             // 无字母
		"Hello 你好",            // 英汉混合
		"程式設計與資料結構",           // 繁体中文
		"程序设计与数据结构",           // 简体中文
	}

	for _, text := range tests {
		fmt.Printf("文本: %q\n", text)
		fmt.Printf("检测到的语言: %s\n\n", detectLanguage(text))
	}
}
