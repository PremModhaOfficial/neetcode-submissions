from typing import List
import math


class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        l, r = 1, max(piles)
        minK = r

        def calculateHr(K: int) -> int:
            sum = 0
            for el in piles:
                sum += math.ceil(el / K)

            return sum

        while l < r:
            mid = (l + r) // 2
            newMin = calculateHr(mid)
            if newMin > h:
                l = mid+1
            else:
                r = mid
            minK = min(mid, minK)

        return l
