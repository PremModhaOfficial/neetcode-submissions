class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        A, B = nums1, nums2
        if len(A) > len(B):
            A, B = B, A

        tl = len(A)+len(B)
        hl = tl//2

        l, r = 0, len(A)-1
        while True:
            i = (l+r)//2  # A
            j = hl - i - 2  # B
            lA = A[i] if i >= 0 else float("-infinity")
            lB = B[j] if j >= 0 else float("-infinity")
            rA = A[i+1] if (i+1) < len(A) else float("infinity")
            rB = B[j+1] if (j+1) < len(B) else float("infinity")

            if lA <= rB and lB <= rA:
                if tl % 2:
                    return min(rA, rB)
                else:
                    return (max(lA, lB) + min(rA, rB))/2
            elif lA > rB:
                r = i - 1
            else:
                l = i+1
