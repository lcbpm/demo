package lang

import (
	"fmt"
	"os/exec"
	"testing"
)

func Test_langdetect(t *testing.T) {
	// 运行 Python 脚本
	cmd := exec.Command("/opt/homebrew/bin/python3", "./script.py")
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		// 输出 Python 脚本的结果
		fmt.Println(string(out))
	}
}
