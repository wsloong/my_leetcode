## 面试题05.替换空格

```
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
示例 1：
    输入：s = "We are happy."
    输出："We%20are%20happy."
限制：
    0 <= s 的长度 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

## 思路
```
一
    可以直接 strings.ReplaceAll(s, " ", "%20")

二
    1，定义一个类型为 []rune 的 slice
    2, 遍历字符串，如果不为空，直接append到slice中，如果为空则 append "%20" 到 slice 中
    3，string([]rune)

    T: O(n)
    S: O(n)
```