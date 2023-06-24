#
# @lc app=leetcode.cn id=235 lang=python3
#
# [235] 二叉搜索树的最近公共祖先
#
# https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
#
# algorithms
# Easy (63.38%)
# Likes:    286
# Dislikes: 0
# Total Accepted:    56.7K
# Total Submissions: 88.8K
# Testcase Example:  '[6,2,8,0,4,7,9,null,null,3,5]\n2\n8'
#
# 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
# 
# 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x
# 的深度尽可能大（一个节点也可以是它自己的祖先）。”
# 
# 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
# 
# 
# 
# 
# 
# 示例 1:
# 
# 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
# 输出: 6 
# 解释: 节点 2 和节点 8 的最近公共祖先是 6。
# 
# 
# 示例 2:
# 
# 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
# 输出: 2
# 解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
# 
# 
# 
# 说明:
# 
# 
# 所有节点的值都是唯一的。
# p、q 为不同节点且均存在于给定的二叉搜索树中。
# 
# 
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        # if not root :
        #     return root
        # if p.val < root.val and q.val < root.val :
        #     return self.lowestCommonAncestor(root.left, p, q)
        # elif p.val > root.val and q.val > root.val :
        #     return self.lowestCommonAncestor(root.right, p, q)
        # else
        #     return root

        while root :
            if p.val < root.val and q.val < root.val :
                root = root.left
            elif p.val > root.val and q.val > root.val :
                root = root.right
            else:
                break

        return root



# @lc code=end

if __name__ == "__main__":
    # [6,2,8,0,4,7,9,null,null,3,5]
    n6 = TreeNode(6)
    n2 = TreeNode(2)
    n8 = TreeNode(8)
    n0 = TreeNode(0)
    n4 = TreeNode(4)
    n7 = TreeNode(7)
    n9 = TreeNode(9)
    n3 = TreeNode(3)
    n5 = TreeNode(5)
    n6.left = n2
    n6.right = n8
    n2.left = n0
    n2.right = n4
    n4.left = n3
    n4.right = n5
    n8.left = n7
    n8.right = n9

    Solution().lowestCommonAncestor(n6, n2, n4)

