#
# @lc app=leetcode id=108 lang=python3
#
# [108] Convert Sorted Array to Binary Search Tree
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> TreeNode:

        def bst(l):
            if l == []:
                return None
            m = len(l)//2
            node = TreeNode(l[m])
            node.left = bst(l[:m])
            node.right = bst(l[m+1:])

            return node

        return bst(nums)

# nums = [1, 2, 3]

# print(nums[:0])
# print(nums[4:])
