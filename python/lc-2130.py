# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def pairSum(self, head: Optional[ListNode]) -> int:
        val_list = []
        n = 0
        cur = head
        while cur is not None:
            val_list.append(cur.val)
            cur = cur.next
            n += 1
        max_val = 0
        for i in range(n // 2):
            max_val = max(max_val, val_list[i] + val_list[n - i - 1])
        return max_val
