package main

import (
	"fmt"
	"strconv"
	"strings"
)

// permute 函数用于生成数字数组的全排列
func permute(nums []int) [][]int {
	var result [][]int
	var current []int
	visited := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		// 当当前排列长度等于输入的长度时，加入结果
		if len(current) == len(nums) {
			// 必须复制一份当前排列，否则会受到后续修改影响
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		// 遍历每个元素，尝试加入当前排列
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}

			// 标记当前元素为已访问
			visited[i] = true
			// 选择当前元素并进入下一步递归
			current = append(current, nums[i])

			backtrack()

			// 回溯，撤销选择
			visited[i] = false
			current = current[:len(current)-1]
		}
	}

	// 调用回溯函数
	backtrack()
	return result
}

func main() {
	var n int
	// 输入数组的大小
	fmt.Scanf("%d", &n)

	// 如果 n 为 0，直接返回空
	if n == 0 {
		return
	}

	testSlice := make([]int, n)
	// 输入数组的元素
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &testSlice[i])
	}

	res := permute(testSlice)

	// 输出每个排列，按要求格式输出
	for i, perm := range res {
		strPerm := []string{}
		for _, num := range perm {
			strPerm = append(strPerm, strconv.Itoa(num))
		}
		// 不对最后一行添加换行符，避免多余换行
		if i == len(res)-1 {
			fmt.Print(strings.Join(strPerm, " ")) // 使用 fmt.Print 避免额外换行
		} else {
			fmt.Println(strings.Join(strPerm, " "))
		}
	}
}
