# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution:
    def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        n = 0
        tmp = head
        while tmp != None:
            tmp = tmp.next
            n += 1
        pre = None
        cur = head
        k = n // 2
        if k == 0:
            return head.next
        for i in range(k):
            pre = cur
            cur = cur.next
        pre.next = cur.next
        return head
