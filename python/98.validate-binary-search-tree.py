#
# @lc app=leetcode id=98 lang=python3
#
# [98] Validate Binary Search Tree
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


# Recursive
# class Solution:
#     def isValidBST(self, root: TreeNode) -> bool:

#         def helper(root, ret):
#             if root == None:
#                 return True

#             isLleftValid = helper(root.left, ret)

#             if len(ret) == 0:
#                 ret.append(root.val)
#             else:
#                 if ret[len(ret) - 1] >= root.val:
#                     return False
#                 ret.append(root.val)

#             isRightValid = helper(root.right, ret)

#             return isLleftValid and isRightValid

#         ret = []

#         return helper(root, ret)

# Iterative


class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        lastValue = float("-inf")
        s = []
        curNode = root
        while curNode != None or s != []:
            while curNode != None:
                s.append(curNode)
                curNode = curNode.left
            curNode = s.pop()
            if lastValue >= curNode.val:
                return False
            else:
                lastValue = curNode.val
            curNode = curNode.right

        return True
