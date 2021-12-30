package main

// 空间复杂度优化的典型案例，大数据中经常用到，如果是开闭很长的biset，还可以优化成compression bitset, 需要tradoff

/* //dp[i][sum] bool 二维表 => bitset
   bool canPartition(vector<int>& nums) {
       int sum = accumulate(nums.begin(), nums.end(), 0);
       bitset<sum>bits(1);
       if(sum & 1) return false; // sum为奇数
       for (int &num : nums) bits |= (bits << num);
       return bits[sum >> 1];
   }
*/

// 降低空间复杂度 和 逻辑剪枝
func canPartitionV4(nums []int) bool {
	var sum, maxNum int
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	part := sum >> 1
	if sum&1 != 0 || len(nums) < 2 || maxNum > part {
		return false
	}
	bytes := make([]byte, part+1)
	bytes[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := part; j >= nums[i]; j-- {
			bytes[j] |= bytes[j-nums[i]]
		}
		//判断中位数如果是1，直接返回true
		if bytes[part] == 1 {
			return true
		}
	}

	return false
}

// 降低空间复杂度 和 逻辑剪枝
func canPartitionV3(nums []int) bool {
	var sum, maxNum int
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	part := sum >> 1
	if sum&1 != 0 || len(nums) < 2 || maxNum > part {
		return false
	}

	bools := make([]bool, part+1)
	bools[0] = true

	for i := 0; i < len(nums); i++ {
		for j := part; j >= nums[i]; j-- {
			bools[j] = bools[j] || bools[j-nums[i]]
		}
		if bools[part] == true {
			return true
		}
	}

	return false
}

// dp[i][j] 表示：在添加第i个数nums[i-1]后，包里面数字的和能否达到j (定义问题是关键)
// dp[i][0] = true 初始 和为0 能到达0
// 迁移状态判断：
// 添加第i个数字nums[i-1]时，
// 1. 第i个数字nums[i-1] < 目标和j，也就是数字太大，则肯定不能达到，为上次d[i-1][j]的状态
// dp[i][j] = dp[i-1][j]
// 2. 第i个数字nums[i-1] = 目标和j，则可到达状态，继续下次状态
// dp[i][j] = true
// 3. 第i个数字nums[i-1] < 目标和j，则有两种情况：
// 一种是将nums[i-1]添加进来， dp[i-1][j-nums[i-1]]
// 另一种是不添加，dp[i-1][j]
// 满足其中一个可到达的迁移状态方程： dp[i][j] = dp[i-1][j-nums[i-1] || dp[i-1][j]
func canPartitionV2(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	sum := 0
	for _, item := range nums {
		sum += item
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if nums[i-1] == j {
				dp[i][j] = true
				continue
			}
			if nums[i-1] < j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[n][target]
}

// 通过子集求解，空间复杂度O(2^n) 过高,耗内存空间
func canPartitionV1(nums []int) bool {
	n := len(nums)
	ans := [][]int{}
	ans = append(ans, []int{})
	for i := 0; i < n; i++ {
		m := len(ans)
		for j := 0; j < m; j++ {
			tmpArr := make([]int, len(ans[j]))
			copy(tmpArr, ans[j])
			ans = append(ans, append(tmpArr, nums[i]))
		}
	}
	//以上是求子集, 然后折叠求和比较

	sum := func(nums []int) int {
		res := 0
		for _, num := range nums {
			res += num
		}
		return res
	}

	//fmt.Println(ans)
	l := len(ans)
	for k := 0; k < l/2; k++ {
		if sum(ans[k]) == sum(ans[l-k-1]) {
			return true
		}
	}

	return false

}

func main() {
	tCases := [][]int{
		{1, 2, 3},
		{1, 2, 13},
		{1, 2, 13, 10},
		//bad case
		{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 99, 97},
	}

	for _, item := range tCases {
		res := canPartitionV4(item)
		println(res)
	}
}
