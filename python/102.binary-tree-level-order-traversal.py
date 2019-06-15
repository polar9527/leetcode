#
# @lc app=leetcode id=102 lang=python3
#
# [102] Binary Tree Level Order Traversal
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def levelOrder(self, root: TreeNode):
        qAll = []
        retAll = []
        if root == None:
            return qAll
        qAll.append([root])
        while len(qAll) != 0:
            nodes = qAll.pop(0)
            retLevel = []
            qNextLevel = []
            for node in nodes:
                retLevel.append(node.val)
                if node.left != None:
                    qNextLevel.append(node.left)
                if node.right != None:
                    qNextLevel.append(node.right)
            if len(retLevel) != 0:
                retAll.append(retLevel)
            if len(qNextLevel) != 0:
                qAll.append(qNextLevel)
        return retAll


# if __name__ == "__main__":
#     root = TreeNode(3)
#     root.left = TreeNode(9)
#     root.right = TreeNode(20)
#     root.right.left = TreeNode(15)
#     root.right.right = TreeNode(7)
#     ret = Solution().levelOrder(root)
#     print(ret)
