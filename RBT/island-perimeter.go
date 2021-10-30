package main

func islandPerimeter(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	l := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				continue
			}
			ll := 4
			if i-1 >= 0 && grid[i-1][j] == 1 {
				ll--
			}
			if i+1 < n && grid[i+1][j] == 1 {
				ll--
			}
			if j-1 >= 0 && grid[i][j-1] == 1 {
				ll--
			}
			if j+1 < m && grid[i][j+1] == 1 {
				ll--
			}
			l += ll
		}
	}
	return l
}

func islandPerimeterV2(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 { //向上，向左 边界
			return 1
		}
		if i >= n || j >= m { //向下，向右 边界
			return 1
		}
		if grid[i][j] == 0 { //截止边界, 有边
			return 1
		}
		if grid[i][j] == 2 { //截止边界，已尽遍历过
			return 0
		}
		grid[i][j] = 2 //标记遍历过
		return dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
	}

	l := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				l += dfs(i, j)
			}
		}
	}

	return l
}

func main() {
	testCases := [][][]int{
		{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0},
		},
	}

	for _, testCase := range testCases {
		max1 := islandPerimeter(testCase)
		max2 := islandPerimeterV2(testCase)
		println(max1, max2)
	}

}
