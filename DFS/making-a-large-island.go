package main

func DFS(i, j, n, m int, grid [][]int) int {
	if i >= n || j >= m || i < 0 || j < 0 {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0

	return DFS(i-1, j, n, m, grid) +
		DFS(i+1, j, n, m, grid) +
		DFS(i, j+1, n, m, grid) +
		DFS(i, j-1, n, m, grid) + 1
}

func maxIsland(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])

	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				res := DFS(i, j, n, m, grid)
				if res > max {
					max = res
				}
			}
		}
	}

	return max
}

func copyTwoDimensionInt(src [][]int) [][]int {
	n := len(src)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, len(src[0]))
		copy(res[i], src[i])
	}
	return res
}

func largestIsland(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cGrid := copyTwoDimensionInt(grid)
			if grid[i][j] == 0 {
				cGrid[i][j] = 1
			}
			//fmt.Println(cGrid)
			res := maxIsland(cGrid)
			if res > max {
				max = res
			}
		}
		//fmt.Println(grid)
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
		max := largestIsland(testCase)
		println(max)
		return
	}

}
