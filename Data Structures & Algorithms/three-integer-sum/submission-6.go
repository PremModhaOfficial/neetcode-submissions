
import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	ans := [][]int{}

	for inx := range nums {
		if inx != 0 && nums[inx-1] == nums[inx] {
			continue
		}

		l, r := inx+1, len(nums)-1

		
		for l < r {
			sum := nums[inx] + nums[l] + nums[r]
 			if sum == 0 {
				ans = append(ans, []int{nums[inx], nums[l], nums[r]})
				l++
				r--
				// skip duplicates
				for l < r && nums[l] == nums[l-1] {
						l++
				}
				for l < r && nums[r] == nums[r+1] {
						r--
				}
			}
			if sum > 0 {
				r--
				continue
			}
			if sum < 0 {
				l++
				continue
			}

		}
	}

	return ans
}
