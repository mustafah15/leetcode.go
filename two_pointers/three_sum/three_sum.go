package two_pointers

import "sort"

func three_sum(nums []int) [][]int {
	var res [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		low, high := i+1, len(nums)-1

		for low < high {
			target := nums[i] + nums[low] + nums[high]

			if target > 0 {
				high--
			} else if target < 0 {
				low++
			} else {
				res = append(res, []int{nums[i], nums[low], nums[high]})

				low++
				for nums[low] == nums[low-1] && low < high {
					low++
				}
			}
		}
	}

	return res
}
