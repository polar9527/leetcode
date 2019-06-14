#
# @lc app=leetcode id=144 lang=python3
#
# [144] Binary Tree Preorder Traversal
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

# Iterative
# class Solution:
#     def preorderTraversal(self, root: TreeNode) -> List[int]:
#         ret = []
#         s = []

#         curNode = root
#         while curNode != None or s != []:
#             while curNode != None:
#                 ret.append(curNode.val)
#                 s.append(curNode.right)
#                 curNode = curNode.left

#             if s == []:
#                 break

#             curNode = s.pop()

#         return ret

# Recursive


class Solution:
    def preorderTraversal(self, root: TreeNode) -> List[int]:

        def helper(root, retList):
            if root == None:
                return

            retList.append(root.val)
            helper(root.left, retList)
            helper(root.right, retList)

        ret = []

        helper(root, ret)
        return ret
