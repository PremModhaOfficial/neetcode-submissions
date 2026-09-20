class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        l, r = 0 , len(matrix)* len(matrix[0])

        while l < r:
            mid = l + (r - l) // 2
            midVal = matrix[mid//len(matrix[0])][mid % len(matrix[0])]

            if midVal > target:
                r = mid
            elif midVal < target:
                l = mid+1
            else:
                return True
            
        return False