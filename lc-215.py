import copy


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        pivot = random.choice(nums)
        left = [val for val in nums if val > pivot]
        mid = [val for val in nums if val == pivot]
        right = [val for val in nums if val < pivot]
        if k <= len(left):
            return self.findKthLargest(left, k)
        elif k > len(left) + len(mid):
            return self.findKthLargest(right, k - len(left) - len(mid))
        else:
            return mid[0]
