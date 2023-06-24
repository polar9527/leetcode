# !/usr/bin/env python3

"""
# Definition for a Node.
class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
"""

class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:

    def treeToDoublyList(self, root: 'Node') -> 'Node':
        pass


class Solution1:

    def __init__(self):
        self.list = []

    def treeToDoublyList(self, root: 'Node') -> 'Node':
        if not root:
            return root

        self.inorder(root)
        for i, node in enumerate(self.list):
            try:
                next_node = self.list[i + 1]
                node.right = next_node
                next_node.left = node
            except IndexError as e:
                self.list[0].left = node
                node.right = self.list[0]
                break
        return self.list[0]

    def inorder(self, root):
        if not root:
            return
        self.inorder(root.left)
        self.visit(root)
        self.inorder(root.right)

    def visit(self, root):
        self.list.append(root)


if __name__ == "__main__":
    n4 = Node(4)
    n2 = Node(2)
    n5 = Node(5)
    n4.left = n2
    n4.right = n5
    n1 = Node(1)
    n3 = Node(3)
    n2.left = n1
    n2.right = n3

    Solution().treeToDoublyList(n4)
