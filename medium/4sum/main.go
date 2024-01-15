package main

func main() {

}

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0, len(nums))
	cache := make(map[int][]int)

	for _, num := range nums {
		want := target - num
		if v, ok := cache[want]; ok {
			v
			res = append(res, cache)
		}
	}

	return res

}
