class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        maxArr = 0
        stack = []

        for i, h in enumerate(heights):
            start = i
            while stack and stack[-1][1] > h:
                index, height = stack.pop()
                maxArr = max(maxArr, height * (i - index))
                start = index

            stack.append((start, h))

        for i, h in stack:
            maxArr = max(maxArr, h * (len(heights)-i))

        return maxArr
