func search(nums []int, target int) int {
    l, r := 0, len(nums)

    mid := (l+r)/2
    for l < r {
        mid = (l+r)/2
        switch {
            case nums[mid] > target:
                r = mid
            case nums[mid] < target:
                l = mid+1
            default:
                return mid
        }
    }
        return -1

}
