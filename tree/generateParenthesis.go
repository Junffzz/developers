package main

func main() {

}

/**
 * #22.有效括号组合
数字n代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。


示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

n=1 ()
n=2 ()(),(())
n=3 ((())),(()()),(())(),()(()),()()()

思路：
1.数学归纳，代码不太好写。
2.递归，dfs深度优先搜索。2*n,但是不是最优  O()
3.改进。a.局部不合法，不再递归。b.左括号用了多少，右括号用了多少
 * @Description: 
 * @Params: 
 * @date: 2021/2/10 
 */
func generateParenthesis(n int) []string {

}