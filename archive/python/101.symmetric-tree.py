#
# @lc app=leetcode id=101 lang=python3
#
# [101] Symmetric Tree
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

# Recursive


# class Solution:
#     def isSymmetric(self, root: TreeNode) -> bool:
#         def isMirror(node1, node2):

#             if node1 == None and node2 == None:
#                 return True
#             if node1 == None or node2 == None:
#                 return False

#             return node1.val == node2.val and \
#                 isMirror(node1.left, node2.right) and \
#                 isMirror(node1.right, node2.left)

#         if root == None:
#             return True

#         return isMirror(root.left, root.right)

# Iterative


class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        q = []
        if root == None:
            return True

        q.append(root.left)
        q.append(root.right)
        while len(q) != 0:
            node1 = q.pop(0)
            node2 = q.pop(0)
            if node1 == None and node2 == None:
                continue
            if node1 == None or node2 == None:
                return False
            if node1.val != node2.val:
                return False
            q.append(node1.left)
            q.append(node2.right)
            q.append(node1.right)
            q.append(node2.left)
        return True
            
