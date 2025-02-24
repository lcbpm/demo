package lang

import (
	"fmt"
	"testing"

	"github.com/abadojack/whatlanggo"
)

func isSimplifiedChinese(text string) bool {
	// 简体中文字符通常包含这些字符：abc, 你我他等
	// 如果包含一些简体汉字就认为是简体中文
	for _, r := range text {
		if r >= 0x4e00 && r <= 0x9fff { // 这是汉字的Unicode范围
			// 通过Unicode范围，简体字通常在一些较高的Unicode区间
			if r == 0x4e00 || r == 0x9fff {
				return true
			}
		}
	}
	return false
}
func Test_what(t *testing.T) {
	// 你想检测的文本
	texts := []string{
		"Lena Blackwood is a woman with long, flowing ebony hair that cascades down Lena Blackwoodr back. SLena Blackwood is often seen in an elegant black dress, exuding an air of sophistication and mystery. Her piercing gaze is accentuated by tLena Blackwood striking red lipstick adorning Lena Blackwoodr lips, giving Lena Blackwoodr a bold and confident demeanor. Against a backdrop of crimson and ivory hues, Lena stands out like a captivating enigma, embodying grace and allure.",
		"張樂天是你哥哥最好的朋友。你穿著運動胸罩和超短裙下樓。整個籃球隊，包括樂天，都盯著你看。", // 英文
		"你好，世界",                 // 中文
		"こんにちは、世界",              // 日文
		"Bonjour tout le monde", // 法语
		"Hola, mundo",           // 西班牙语
	}

	// 遍历每个文本并检测语言
	for _, text := range texts {
		lang := whatlanggo.Detect(text) // 检测文本的语言及置信度
		fmt.Printf("Text: %s\nDetected Language: %v\n", text, lang.Lang.Iso6391())
		if lang.Lang.Iso6391() == "zh" { // 检查是否是中文
			// 进一步判断简体或繁体
			if isSimplifiedChinese(text) {
				fmt.Printf("Text: %s\nDetected Language: Chinese (Simplified)\n", text)
			} else {
				fmt.Printf("Text: %s\nDetected Language: Chinese (Traditional)\n", text)
			}
		} else {
			fmt.Printf("Text: %s\nDetected Language: %v\n", text, lang.Lang.Iso6391())
		}
	}
}
