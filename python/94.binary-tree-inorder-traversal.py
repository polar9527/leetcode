#
# @lc app=leetcode id=94 lang=python3
#
# [94] Binary Tree Inorder Traversal
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

# Recursive
# class Solution:
#     def inorderTraversal(self, root: TreeNode) -> List[int]:

#         def helper(root: TreeNode, retList: List[int]):
#             if root == None:
#                 return

#             if root.left != None:
#                 helper(root.left, retList)

#             retList.append(root.val)

#             if root.right != None:
#                 helper(root.right, retList)

#         ret = []
#         helper(root, ret)

#         return ret

# Iterative using stack
class Solution:
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        ret = []
        s = []

        curNode = root
        while curNode != None or s != []:
            while curNode != None:
                s.append(curNode)
                curNode = curNode.left

            curNode = s.pop()
            ret.append(curNode.val)
            curNode = curNode.right

        return ret
