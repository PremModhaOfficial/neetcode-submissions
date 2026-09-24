class Solution:
    def findMin(self, nums: List[int]) -> int:
        minF = nums[0]

        l, r = 0, len(nums)
        minInd = 0

        while l < r:
            mid = (l+r)//2

            if minF > nums[mid]:
                minInd = mid
                r = mid
            else:
                l = mid+1

        return nums[minInd]





        