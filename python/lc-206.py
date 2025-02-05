# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None:
            return head
        p1, p2 = head, head.next
        p3 = None
        if p2 is not None:
            p3 = p2.next
        while p2 is not None:
            p2.next = p1
            p1 = p2
            p2 = p3
            if p3 is not None:
                p3 = p3.next
        head.next = None
        return p1
