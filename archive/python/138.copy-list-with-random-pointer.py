#
# @lc app=leetcode id=138 lang=python3
#
# [138] Copy List with Random Pointer
#

# @lc code=start
"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""

class Solution:
    def copyRandomList(self, head: 'Node') -> 'Node':
        if head == None :
            return head

        hashTable = {}

        n = head
        cPre = cHead = Node(n.val)
        hashTable[n] = cHead
        n = n.next
        while n :
            copyNode = Node(n.val)
            cPre.next = copyNode
            cPre = cPre.next

            hashTable[n] = copyNode
            n = n.next

        n = head
        while n:
            if n.random != None :
                hashTable[n].random = hashTable[n.random]
            n = n.next

        return cHead
        
# @lc code=end

