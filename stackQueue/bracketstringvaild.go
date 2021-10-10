/*
@Time : 2020/12/23 13:31
@Author : ZhaoJunfeng
@File : bracketstringvaild
*/
package main

/**
string 大、中、小括号=》合法？
1."()"
2."()[]"
3."([)]"
4."((([])))"
5."]][["

解题思路：堆栈
a.左=>push
b.右=>pop/peek
c.stack empty

时间复杂度：O(1)*n=>O(n)
空间复杂度：O(n)
*/

type Stack []string

func (s *Stack) Pop() string {
	ns := *s
	v := ns[len(ns)-1]
	*s = ns[:len(ns)-1]

	return v
}

func IsValid(s string) bool {
	var stack Stack
	hash_map := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, b := range s {
		if _, ok := hash_map[string(b)]; !ok {
			stack = append(stack, string(b))
		} else if len(stack) == 0 || stack.Pop() != hash_map[string(b)] {
			return false
		}
	}
	return len(stack) == 0
}

