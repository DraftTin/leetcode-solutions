# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def oddEvenList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        odd_head = ListNode()
        odd_cur = odd_head
        even_head = ListNode()
        even_cur = even_head
        i = 0
        while head is not None:
            if i % 2 == 0:
                odd_cur.next = head
                odd_cur = odd_cur.next
            else:
                even_cur.next = head
                even_cur = even_cur.next
            i += 1
            head = head.next

        even_cur.next = None
        odd_cur.next = even_head.next
        return odd_head.next
