package algorithms_go

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)


/**
	Desc： 对某一字符串排序
*/


// 方法一：首先,字符串转为[]byte, 对切片排序, 最后切片转字符串
func SortString2(str string) {
	// str := "eat"
	b := []byte(str)		// string转换成字节  [101 97 116]
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	str = string(b)			// 字节转字符串
}


// 方法二：首先字符串分割,得到字符串切片, 对切片排序, strings.join连接成新的字符串
func SortString(str string) string {
	// str := "eat"
	split := strings.Split(str, "")	// 字符串切片 []string[e a t]
	sort.Strings(split)						// 升序排序   []string[a e t]
	return strings.Join(split, "")		// 切片转化为字符串 aet
}




/**
	Desc： 字符串查找测试
*/


// 前缀测试
func HasPrefix(s, prefix string) bool  {
	return len(s) > len(prefix) && s[:len(prefix)] == prefix
}

// 后缀测试
func HasSuffix(s, suffix string) bool  {
	return len(s) > len(suffix) && s[len(s) - len(suffix):] == suffix
}

// 子串测试
func Contains(s, substr string) bool  {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}

	return false
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')	// 添加任意字符的UTF8编码

	return buf.String()
}

// Output：fmt.Println(intsToString([]int{1,2,3}))   [1, 2, 3]


