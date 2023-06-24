#
# @lc app=leetcode.cn id=297 lang=python3
#
# [297] 二叉树的序列化与反序列化
#
# https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/description/
#
# algorithms
# Hard (46.77%)
# Likes:    191
# Dislikes: 0
# Total Accepted:    23.6K
# Total Submissions: 49.9K
# Testcase Example:  '[1,2,3,null,null,4,5]'
#
# 
# 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
# 
# 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 /
# 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
# 
# 示例: 
# 
# 你可以将以下二叉树：
# 
# ⁠   1
# ⁠  / \
# ⁠ 2   3
# ⁠    / \
# ⁠   4   5
# 
# 序列化为 "[1,2,3,null,null,4,5]"
# 
# 提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode
# 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
# 
# 说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
# 
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


import random


class Codec:

    def __init__(self):
        self.res = []
        self.flag = random.choice(["dfs", "bfs"])

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        switch = {"dfs": self.dfsSerialize,
                  "bfs": self.bfsSerialize}

        return switch[self.flag](root)

    def bfsSerialize(self, root):
        if not root:
            return ""

        queue = [root]
        res = []
        while len(queue) != 0:
            cur_node = queue[0]
            if not cur_node:
                res.append(str(cur_node.val))
                queue.extend([cur_node.left, cur_node.right])
            else:
                res.append("null")
            queue = queue[1:]

        return ",".join(res)

    def dfsSerialize(self, root):
        if not root:
            return "null"
        val = str(root.val)
        return ",".join([val, self.bfsSerialize(root.left), self.bfsSerialize(root.right)])

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """

        switch = {"dfs": self.dfsDeserialize,
                  "bfs": self.bfsDeserialize}

        return switch[self.flag](data)

    def bfsDeserialize(self, data):
        if len(data) == 0:
            return None
        queue = []
        res = data.split(",")
        root = TreeNode(res[0])
        queue.append(root)
        res = res[1:]

        while len(queue) != 0 and len(res) > 1:
            cur_node = queue[0]
            if res[0] != "null":
                cur_node.left = TreeNode(int(res[0]))
                queue.append(cur_node.left)

            if res[1] != "null":
                cur_node.right = TreeNode(int(res[1]))
                queue.append(cur_node.right)
            res = res[2:]
            queue = queue[1:]

        return root

    def dfsDeserialize(self, data):
        if len(data) == 0:
            return None
        self.res = data.split(",")
        return self.dfsDeserializeCore()

    def dfsDeserializeCore(self):
        node_val = self.res[0]
        self.res = self.res[1:]
        if node_val == "null":
            return None

        root = TreeNode(int(node_val))
        root.left = self.dfsDeserializeCore()
        root.right = self.dfsDeserializeCore()
        return root



# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))
# @lc code=end

if __name__ == "__main__":
    data = "1,2,3,null,null,4,5"
    s = Codec()
    root = s.bfsDeserialize(data)
    print(s.bfsSerialize(root))