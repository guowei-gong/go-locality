package go_locality

// Loop 访问内存, 并写入值
// 通过 step 定制每次遍历的跨度
func Loop(nums []int, step int) {
	l := len(nums)
	for i := 0; i < step; i++ {
		for j := i; j < l; j += step {
			nums[j] = 4
		}
	}
}
