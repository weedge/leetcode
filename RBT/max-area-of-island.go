package main

// 深度优先(多叉树🌲深度优先遍历)
func maxAreaOfIsland(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= n || j >= m {
			return 0
		}
		if i < 0 || j < 0 {
			return 0
		}
		if grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = 0
		return dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1) + 1
	}

	max := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] != 1 {
				continue
			}
			res := dfs(i, j)
			if res > max {
				max = res
			}
		}
	}

	return max
}

func main() {
	testCases := [][][]int{
		{
			{1, 0, 1, 1, 1},
			{1, 0, 1, 0, 1},
			{1, 1, 1, 0, 1},
		},
		{
			{1, 1, 1},
			{0, 1, 0},
			{1, 1, 1},
		},
		{
			{1, 1},
		},
		{
			{0, 1},
			{1, 0},
		},
		{
			{1, 1, 1, 0},
			{1, 1, 0, 0},
			{1, 0, 0, 0},
		},
		{
			{1, 1, 1, 0},
			{1, 1, 0, 0},
			{1, 0, 1, 0},
			{1, 0, 0, 1},
		},
	}

	for _, testCase := range testCases {
		max := maxAreaOfIsland(testCase)
		println(max)
	}

}
