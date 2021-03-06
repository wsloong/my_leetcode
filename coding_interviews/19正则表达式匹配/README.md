### 19正则表达式匹配

```
请实现一个函数用来匹配包含'. '和'*'的正则表达式。
模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。
在本题中，匹配是指字符串的所有字符匹配整个模式。
例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

示例 1:
    输入:
        s = "aa"
        p = "a"
    输出: false
    解释: "a" 无法匹配 "aa" 整个字符串。

示例 2:
    输入:
        s = "aa"
        p = "a*"
    输出: true
    解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

示例 3:
    输入:
        s = "ab"
        p = ".*"
    输出: true
    解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

示例 4:
    输入:
        s = "aab"
        p = "c*a*b"
    输出: true
    解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。

示例 5:
    输入:
        s = "mississippi"
        p = "mis*is*p*."
    输出: false
    s 可能为空，且只包含从 a-z 的小写字母。
    p 可能为空，且只包含从 a-z 的小写字母以及字符 . 和 *，无连续的 '*'。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

#### 思路

手抄于[官方思路](https://leetcode-cn.com/problems/regular-expression-matching/solution/zheng-ze-biao-da-shi-pi-pei-by-leetcode-solution/#comment)

方法： 动态规划

题目中匹配是一个 `逐步匹配` 的过程:每次从字符串 `p` 中取出 `一个字符` 或者 `字符+星号` 的组合，并在 `s` 中进行匹配。
对于 `p `中 `一个字符` ：它只能在 `s` 中匹配一个字符，匹配的方法具有唯一性；
对于 `p` 中 `字符+星号` ：它可以在 `s` 中匹配任意自然数个字符，并不具有唯一性，可以考虑使用动态规划，对匹配方案进行枚举

使用 `f[i][j]` 标示 `s` 的前 `i` 个字符和 `p` 中前 `j` 个字符是否能够匹配，在进行状态转移时，我们考虑 `p` 的第 `j` 个字符匹配情况：

* 如果 `p` 的第 `j` 个字符是小写字母，那么我们必须在`s`中匹配一个相同的小写字母，即

```
                / f[i-1][j-1], s[i]=p[j]
    f[i][j] =   |
                \ false, s[i] != p[j]
```

也就是说，如果 `s` 的第 `i` 个字符与 `p` 的第 `j` 个字符不同，那么无法进行匹配；
否则可以匹配两个字符串的最后一个字符，完整匹配结果取决于两个字符串前面的部分。


* 如果 `p` 的第 `j` 个字符是 `*`,那么就表示我们可以对 `p` 第 `j-1` 个字符匹配任意自然数次。
在匹配 `0` 次的情况下，我们有 `f[i][j] = f[i][j-2]` ，也就是我们 `浪费` 了一个 `字符+星号` 的组合，没有匹配任何 `s` 中的字符。
在匹配 `1,2,3....` 次的情况下，类似地我们有

```
f[i][j] = f[i-1][j-2], if s[i] = p[j-1]
f[i][j] = f[i-2][j-2], if s[i-1] = s[i] = p[j-1]
f[i]f[j] = f[i-3][j-2], if s[i-2] = s[i-1] = s[i] = p[j-1]
......

```

如果我们通过这种方式进行转移，那么我们就需要枚举这个组合到底匹配了 `s` 中的几个字符，会导致时间复杂度增加，并且代码编写起来十分麻烦。
我们不妨换个角度考虑这个问题，`字符+星号` 的组合在匹配过程中，本质上只会有两种情况:

1. 匹配 `s` 末尾的一个字符，将该字符扔掉，而该组合可以继续匹配
2. 不匹配字符，将该组合扔掉，不进行匹配

如果按照这个角度进行思考，我们就可以写出很精巧的状态转移方程:

```
                / f[i-1][j] or f[i][j-2],  s[i]=p[j-1]
    f[i][j] =   |
                \ f[i][j-2],            s[i] != p[j-1]
```

* 在任意情况下，只要 `p[j]` 是 `.` ，那么 `p[j]` 一定成功匹配 `s` 中的任意一个小写字母

最终的状态转移方程如下：

```
                                         
                                        / f[i-1][j-1],  matches(s[i], p[j])
                / if (p[j] != '*') =    |
                |                       \ false,  otherwise
    f[i][j] =   |
                |               / f[i-1][j] or f[i][j-2],  matches(s[i], p[j-1])
                \ otherwise =   |
                                \ f[i][j-2], otherwise
```
其中 `matches(x,y)` 判断两个字符是否匹配的辅助函数。只有当 `y` 是 `.` 或者 `x` 和 `y` 本身相同时，这两个字符才会匹配

`T: O(mn)`, 其中 `m` 和 `n` 分别是字符串 `s` 和 `p` 的长度。我们需要计算出所有的状态，并且每个状态在进行转移时的时间复杂度为 `O(1)`
`S: O(mn)`, 存储所有状态使用的空间