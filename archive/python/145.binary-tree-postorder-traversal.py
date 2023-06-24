#
# @lc app=leetcode id=145 lang=python3
#
# [145] Binary Tree Postorder Traversal
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

# Iterative


# class Solution:
#     def postorderTraversal(self, root: TreeNode) -> List[int]:

#         def gotoHLVL(s):
#             if len(s) == 0:
#                 return
#             while s[len(s)-1][0] != None:
#                 node = s[len(s)-1][0]
#                 if node.left != None:
#                     if node.right != None:
#                         s.append((node.right, node, None))
#                     s.append((node.left, node, node.right))
#                 else:
#                     s.append((node.right, node, None))
#             s.pop()

#         ret = []
#         s = []
#         parentNode = None
#         rSibling = None
#         if(root != None):
#             s.append((root, parentNode, rSibling))
#         curNode = root
#         while len(s) != 0:
#             if s[len(s) - 1][0] != parentNode:
#                 gotoHLVL(s)

#             curNode, parentNode, _ = s.pop()
#             ret.append(curNode.val)

#         return ret

# Recursive


class Solution:
    def postorderTraversal(self, root: TreeNode) -> List[int]:
        def helper(root, ret):
            if root == None:
                return None

            helper(root.left, ret)
            helper(root.right, ret)
            ret.append(root.val)

        ret = []
        helper(root, ret)

        return ret
