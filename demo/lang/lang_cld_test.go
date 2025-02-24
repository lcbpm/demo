package lang

//
//import (
//	"fmt"
//	"log"
//	"testing"
//
//	"github.com/jmhodges/gocld3/cld3"
//)
//
//func Test_cld(t *testing.T) {
//	// 创建一个 NNetLanguageIdentifier（该类型是 gocld3 提供的语言检测器）
//	identifier, err := cld3.NewLanguageIdentifier(0, 512)
//	if err != nil {
//
//	}
//	defer cld3.FreeLanguageIdentifier(identifier) // 确保检测器关闭，释放资源
//
//	// 待检测的文本
//	texts := []string{
//		"你好，世界",            // 中文
//		"Hello, world!",    // 英文
//		"こんにちは、世界",         // 日文
//		"Bonjour le monde", // 法语
//		"Привет, мир",      // 俄语
//		"Hola, mundo",      // 西班牙语
//	}
//
//	// 遍历检测文本
//	for _, text := range texts {
//		lang, err := cldDetectLanguage(identifier, text)
//		if err != nil {
//			log.Println("Error detecting language:", err)
//			continue
//		}
//		fmt.Printf("Text: %s\nDetected Language: %s\n\n", text, lang)
//	}
//}
//
//// detectLanguage 使用 gocld3 进行语言检测
//func cldDetectLanguage(detector cld3.LanguageIdentifier, text string) (string, error) {
//	// 使用 FindLanguage 方法检测文本语言
//	result := detector.FindLanguage(text)
//
//	// 返回检测到的语言代码 (ISO 639-1)
//	return result.Language, nil
//}
