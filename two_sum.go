package main

func twoSum(nums []int, target int) [2]int {
	hashMap := map[int]int{}

	for i:=0; i<len(nums); i++ {
		delta := target - nums[i]

		idx, persistInHash := hashMap[delta]

		if persistInHash {
			return [2]int{i, idx}
		} else {
			hashMap[nums[i]] = i
		}
	}

	return [2]int{}
}
