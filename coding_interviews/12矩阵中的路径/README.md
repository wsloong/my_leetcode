### 12矩阵中的路径
```
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。
如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。

例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
 
示例 1：
    输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
    输出：true

示例 2：
    输入：board = [["a","b"],["c","d"]], word = "abcd"
    输出：false
提示：
1 <= board.length <= 200
1 <= board[i].length <= 200

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

### 思路

* 以下内容摘自 [jyd大佬的解题思路](https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/solution/mian-shi-ti-12-ju-zhen-zhong-de-lu-jing-shen-du-yo/)

#### 原理
* 深度优先算法：通过DFS递归，先朝一个方向搜索到底，再回溯至上一个节点，沿另一个方向搜索，依次类推
* 剪枝：在搜索中，遇到`这条路不可能或这目标字符串匹配成功`的情况，则应立即返回

#### 刨析

* 递归参数：当前元素在矩阵`board`中的行列索引`i`和`j`,当前目标字符在`word`中索引`k`
* 终止条件:
    * 行或者列索引越界
    * 当前矩阵元素与目录字符不同
    * 当前矩阵元素已经访问过
    * 字符串`word`已经全部匹配；返回`true`
* 递推工作
    * 标记当前矩阵元素：将`board[i][j]`值暂存于`tmp`,并修改为`/`,代表元素已经访问过了，防止之后搜索是重复访问
    * 搜索下一个单元格，朝当前元素的上下左右四个方向开启下层递归
    * 还原当前矩阵元素: 将`tmp`暂存值还原至`board[i][j]`元素

作者：jyd
链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/solution/mian-shi-ti-12-ju-zhen-zhong-de-lu-jing-shen-du-yo/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。