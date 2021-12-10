package algorithms_go

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"time"
	"fmt"
)

/*
问题描述：
	???

思路解析：
    1.首先对输入参数进行正确分组；
    2.通过集合Map进行去重；
    3.最后对集合Map按照key值排序
回顾总结：
	???

*/
func paixu() {
	counts := make(map[int]int)

	// 输入数据流
	input := bufio.NewScanner(os.Stdin)
	inputslice := []int{}
	c1 := make(chan int, 1)
	fg := false
	fmt.Printf("Please type in something:\n")

	go func() {
		for input.Scan() {
			line := input.Text()
			numline, _ := strconv.Atoi(line)	// string转int
			inputslice = append(inputslice, numline)
			c1 <-numline
		}
	}()

	for  {
		select {
		case <-c1:

		case <-time.After(time.Second * 5):	// 当没有新的输入时（超时5妙），说明输入结束
			fg = true
			goto end
		}
	end:
		if fg {
			break
		} else {
			continue
		}
	}

	// Test数据 fmt.Println(inputslice)
	// Test数据 inputslice := []int{3,2,2,1,11,10,20,40,32,67,40,20,89,300,400,15}

	deadline := inputslice[0]
	strlen := len(inputslice)
	flag := false

	// 数据分组&Map去重
	for i := 1; i < strlen; i++ {
		if (flag) {
			flag = false	// 跳跃非有效数值
			continue
		}
		counts[inputslice[i]]++	// 统计&去重
		if ( i == strlen-1) {
			break			// 防止切片越界
		}

		if (i == deadline) {
			deadline += inputslice[i+1]+1	// 重置每组数据的截止索引
			flag = true
			continue
		}
	}

	// 排序
	keys := []int{}
	for k, _ := range  counts {
		keys = append(keys, k)	// 得到各个key
	}
	sort.Sort(sort.IntSlice(keys))	// IntSlice：	Sort：排序

	for _, v := range keys {
		fmt.Println(v)
	}

}

/**
产生一个int数组，长度为100，并向其中随机插入1-100，并且不能重复

初步实现：
namespace Wolfy.RandomDemo
{
    class Program
    {
        static void Main(string[] args)
        {
            List<int> lst = new List<int>();
            Random r = new Random();
            while (true)
            {
                int temp = r.Next(1, 101);
                if (lst.Count == 100)
                {
                    break;
                }
                if (!lst.Contains(temp))
                {
                    lst.Add(temp);
                }
            }
            for (int i = 0; i < lst.Count; i++)
            {
                Console.WriteLine(lst[i]);
            }
            Console.Read();
        }
    }
}

虽然上面的代码，实现题目的要求，但是如果是1到100万或者更大，这样的每次判断是否包含这样的一个数，势必会影响到性能。

网上找到一种更好的实现方式：

(1)把N个数放到容器A(int数组)中.

(2)从N个数中随机取出1个数放入容器B(int数组)中.

(3)把容器A中最后一个数与随机抽取的数对调 或者 把容器A中最后一个数覆盖随机抽取出来的数.

(4)这时从容器A(假设N个数,索引0 到 索引N-2)之间随机取一个数.再放入容器B中,重复此步骤.

说明:也就是第二次是从容器A中 第一个元素到倒数第二个元素 中随机取一个数.

这种好处是,随机数所取范围逐步缩小,而且杜绝了大数据时集合执行删除操作时产生的瓶颈.

namespace Wolfy.RandomDemo
{
    class Program
    {
        static void Main(string[] args)
        {
            int[] result = GetRandom(100);
            for (int i = 0; i < result.Length; i++)
            {
                Console.WriteLine(result[i]);
            }
            Console.WriteLine("over:" + result.Length);
            Console.Read();
        }
        /// <summary>
        /// 获得无重复随机数组
        /// </summary>
        /// <param name="n">上限n</param>
        /// <returns>返回随机数组</returns>
        static int[] GetRandom(int n)
        {
            //容器A和B
            int[] arryA = new int[n];
            int[] arryB = new int[n];
            //填充容器a
            for (int i = 0; i < arryA.Length; i++)
            {
                arryA[i] = i + 1;
            }
            //随机对象
            Random r = new Random();
            //最后一个元素的索引 如n=100，end=99
            int end = n - 1;
            for (int i = 0; i < n; i++)
            {
                //生成随机数 因为随机的是索引 所以从0到100取，end=100
                //一个大于等于 minValue 且小于 maxValue 的 32 位带符号整数，即：返回的值范围包括 minValue 但不包括 maxValue。
                //如果 minValue 等于 maxValue，则返回 minValue
                //
                int minValue = 0;
                int maxValue = end + 1;
                int ranIndex = r.Next(minValue, maxValue);
                //把随机数放在容器B中
                arryB[i] = arryA[ranIndex];
                //用最后一个元素覆盖取出的元素
                arryA[ranIndex] = arryA[end];
                //缩减随机数生成的范围
                end--;
            }
            //返回生成的随机数组
            return arryB;
        }
    }
}


总结：
	实现方式有很多种，但是如果能用高效的方式就用高效的方式实现。这种生成无重复的随机数，可以在运用在抽奖系统中。

*/


/**
字典树
题目：
给你100000个长度不超过10的单词。对于每一个单词，我们要判断他出没出现过，如果出现了，求第一次出现在第几个位置。 
分析：
如果我们使用一般的方法，每查询一个单词都去遍历一遍，那么时间复杂度将为O(n^2),这对于100000这么大的数据是不能够接受的。假如我们要查找单词student。那我们通过前缀树只需要查找s开头的即可，然后接下来查询t开头的即可，对于大量的数据可以省去不小的时间。
树结构：
count：表示以当前单词结尾的单词数量。
prefix：表示以该处节点之前的字符串为前缀的单词数量。
public class TrieNode {
int count;
int prefix;
TrieNode[] nextNode=new TrieNode[26];
public TrieNode(){
count=0;
prefix=0;
}
}


1. 前缀树的创建
好比假设有b，abc，abd，bcd，abcd，efg，hii 这6个单词,那我们创建trie树就得到

//插入一个新单词
public static void insert(TrieNode root,String str){
if(root==null||str.length()==0){
return;
}
char[] c=str.toCharArray();
for(int i=0;i<str.length();i++){
//如果该分支不存在，创建一个新节点
if(root.nextNode[c[i]-'a']==null){
root.nextNode[c[i]-'a']=new TrieNode();
}
root=root.nextNode[c[i]-'a'];
root.prefix++;//注意，应该加在后面
}

//以该节点结尾的单词数+1
root.count++;
}

2. 查询以str开头的单词数量，查询单词str的数量
//查找该单词是否存在，如果存在返回数量，不存在返回-1
public static int search(TrieNode root,String str){
if(root==null||str.length()==0){
return -1;
}
char[] c=str.toCharArray();
for(int i=0;i<str.length();i++){
//如果该分支不存在，表名该单词不存在
if(root.nextNode[c[i]-'a']==null){
return -1;
}
//如果存在，则继续向下遍历
root=root.nextNode[c[i]-'a'];
}

//如果count==0,也说明该单词不存在
if(root.count==0){
return -1;
}
return root.count;
}

//查询以str为前缀的单词数量
public static int searchPrefix(TrieNode root,String str){
if(root==null||str.length()==0){
return -1;
}
char[] c=str.toCharArray();
for(int i=0;i<str.length();i++){
//如果该分支不存在，表名该单词不存在
if(root.nextNode[c[i]-'a']==null){
return -1;
}
//如果存在，则继续向下遍历
root=root.nextNode[c[i]-'a'];
}
return root.prefix;
}
3.在主函数中测试

public static void main(String[] args){
TrieNode newNode=new TrieNode();
insert(newNode,"hello");
insert(newNode,"hello");
insert(newNode,"hello");
insert(newNode,"helloworld");
System.out.println(search(newNode,"hello"));
System.out.println(searchPrefix(newNode,"he"));
}

输出：3   4

*/


/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母都恰好只用一次。


示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]

测试用例
输入：strs = ["aple","pale","ape","ae","apel","aplle","alple","alpe"]
输出：[["aple","apel","alpe","pale"], ["ape"], ["ae"], ["aplle", "alple"]]

提示：

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母


// 字符串排序
func SortString(str string) string {
	split := strings.Split(str, "")
	sort.Strings(split)

	return strings.Join(split, "")

}

// 相等长度的字符串按照字母异位词分组
func SortGroup(str []string) [][]string {
	// 前置条件判断：切片长度小于一直接返回
	if (len(str) == 1) {
		return [][]string{str}
	}

	// 字符串排序
	sortstr := []string{}
	for _, v := range str {
		sortstr = append(sortstr, SortString(v))
	}

	lenstr := len(str)
	retgroup := [][]string{}
	for i:=0; i<lenstr; i++ {
		if (str[i] != "") {
			tmp := []string{str[i]}
			for j:=i+1; j<lenstr; j++ {
				// 等于则进行分组，并从数组中重置;否则，放到下一轮的循环比较
				if(sortstr[i] == sortstr[j]) {
					tmp = append(tmp, str[j])
					str[j] = ""
				}
			}
			retgroup = append(retgroup, tmp)
		}
	}

	return retgroup
}

func main() {
	//输入：strs = ["aple","pale","ape","ae","apel","aplle","alple","alpe"]
	//输出：[["aple","apel","alpe","pale"], ["ape"], ["ae"], ["aplle", "alple"]]

	input := []string{"aple","pale","ape","ae","apel","apee","aplle","alple","alpe"}

	// 按照长度分组
	inputgroup := make(map[int][]string)
	for _, v := range input {
		inputgroup[len(v)] = append(inputgroup[len(v)], v)
	}

	ret := [][]string{}
	// 字母异位词 辨认
	for _, v := range inputgroup {
		tmpret := SortGroup(v)
		for _, vv := range tmpret {
			ret = append(ret, vv)
		}
	}

	fmt.Println(ret)
}

思考总结：
	1、先对问题分类，宏观上把大问题分解成小问题；
	2、把小问题进行编码实现，需要注意安置条件或者边界条件
	3、先求小，再求大而全：先实现程序，然后再求进一步的改进：内存、时间复杂度、空间负责度
	4、使用切片的时候，切记不要越界。
*/

