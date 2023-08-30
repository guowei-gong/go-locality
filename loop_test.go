package go_locality

import "testing"

// CreateSource 初始化一个数组, 并赋值
func CreateSource(len int) []int {
	nums := make([]int, 0, len)

	for i := 0; i < len; i++ {
		nums = append(nums, i)
	}

	return nums
}

func BenchmarkLoopStep1(b *testing.B) {
	src := CreateSource(10000)

	// 重置计时, 去掉 CreateSource 小号的时间
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 1)
	}
}

func BenchmarkLoopStep2(b *testing.B) {
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 2)
	}
}

func BenchmarkLoopStep3(b *testing.B) {
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 3)
	}
}

func BenchmarkLoopStep4(b *testing.B) {
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 4)
	}
}

func BenchmarkLoopStep5(b *testing.B) {
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 5)
	}
}

func BenchmarkLoopStep6(b *testing.B) {
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 6)
	}
}

func BenchmarkLoopStep12(b *testing.B) {
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 12)
	}
}

func BenchmarkLoopStep16(b *testing.B) {
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Loop(src, 16)
	}
}
