package algorithms_go

import (
	"bytes"
	"fmt"
)

/**
	Desc：位基本运算
 */

func operation() {
	var x uint8 = 1<<1 | 1<<5	// 左移：把数值1的二进位分别向左移动1位和5位，然后再取或操作
	var y uint8 = 1<<1 | 1<<2

	// %b:打印二进制格式数字；08:打印至少8个字符宽度
	fmt.Printf("%08b\n",x)		// 00100010  	the set{1, 5}
	fmt.Printf("%08b\n",y)		// 00000110		the set{1, 2}

	fmt.Printf("%08b\n", x&y)	// 与运算 	00000010 the intersection {1}
	fmt.Printf("%08b\n", x|y)	// 或运算	00100110 the union	{1, 2, 5}
	fmt.Printf("%08b\n", x^y)	// 异或运算	00100100 the symetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y)	// 位清空	00100000 the differenct {5}


	for i:=uint(0); i<8; i++ {
		if x&(1<<i) != 0 {		// membership test
			fmt.Println(i)		// 打印x变量二进制中1的位置，也就是集合中的值，"1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1)	// 01000100 the set {2, 6}  注：位运算不会更改目标值。
	fmt.Printf("%08b\n", x>>1)	// 00010001 the set {0, 4}
}


/**
	Desc：未彻底搞明白
*/

type IntSet struct {
	words []uint64
}

// 因为每一个字都有64个二进制位，x/64的商作为字的下标，用x%64的余数作为这个字内的bit所在位置
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}


func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] = 1 << bit
}

// 取并集
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword		// "或"逻辑操作符号|来一次完成64个元素的或计算
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// 将IntSet作为一个字符串来打印。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {			// uint(32位或者64位)
				if buf.Len() > len("{") {
					buf.WriteByte('}')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)	// fmt会直接调用用户定义的String方法
			}
		}
	}

	buf.WriteByte('}')
	return buf.String()
}

func test_one() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())
	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))
}



