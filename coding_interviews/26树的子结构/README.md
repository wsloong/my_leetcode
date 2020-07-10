### 26树的子结构

```
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:

给定的树 A:

     3
    / \
   4   5
  / \
 1   2

给定的树 B：
   4 
  /
 1

返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：
    输入：A = [1,2,3], B = [3,1]
    输出：false

示例 2：
    输入：A = [3,4,5,1,2], B = [4,1]
    输出：true

限制：
0 <= 节点个数 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

#### 思路

若树`B`是树`A`的子结构,则子结构的根节点可能为树A的任意一个节点.因此,判断树B是否是树A的子结构,需完成一下2个步骤
1. 先遍历树`A`的每个节点`N`;(对应`isSubStructure(A.Left,B)`和`isSubStructure(A.Right,B)`)
2. 判断树`A`中`以N为根节点的子树`是否包含树`B`(对应`recur(A,B)`)

##### 算法流程

* recur(A, B):
    1. 终止条件:
        * 节点`B`为空: 说明`B`已经匹配完成,返回`true`;
        * 节点`A`为空: 说明已经越过了树`A`叶子节点,即匹配失败,返回`false`
        * `A`和`B`的值不同:说明匹配失败,返回`false`;
    2. 返回值
        * 判断`A`和`B`的左子树是否相等,即`recur(A.Left, B.Left)`;
        * 判断`A`和`B`的右子树是否相等,即`recur(A.Right, B.Right)`;
* isSubStructure(A, B)函数:
    1. 特例处理: 当`A`或者`B`为空时候,返回`false`;
    2. 返回值: 若树`B`是树`A`的子结构,必须满足一下三种情况之一,因此用 或 `||`链接
        1. 以`节点A为根节点的子树`包含树`B`, 对应`recur(A, B)`
        2. 树`B`是`树A左子树`的子结构, 对应`isSubStructure(A.Left, B)`;
        3. 树`B`是`树A右子树`的子结构, 对应`isSubStructure(A.Right, B)`

#### 时间复杂度和空间复杂度

* T: `O(MN)`; `M`,`N`分别为`树A`和`树B`的节点数量; 先序遍历树 `A` 占用 `O(M)`，每次调用 `recur(A, B)` 判断占用 `O(N)`.
* S: `O(M)`; 当树 `A` 和树 `B` 都退化为链表时，递归调用深度最大。当 `M≤N` 时，遍历树 `A` 与递归判断的总递归深度为 `M` ；当 `M>N` 时，最差情况为遍历至树 `A`叶子节点，此时总递归深度为 `M`

作者：jyd
链接：https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/solution/mian-shi-ti-26-shu-de-zi-jie-gou-xian-xu-bian-li-p/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。