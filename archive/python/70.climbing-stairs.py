#
# @lc app=leetcode id=70 lang=python3
#
# [70] Climbing Stairs
#


class Solution:
    def climbStairs(self, n: int) -> int:
        d = {}

        def helper(n, d):
            if n == 2:
                d[n] = 2
            if n == 1:
                d[n] = 1
            if n not in d.keys():
                d[n] = helper(n - 1, d) + helper(n - 2, d)

            return d[n]

        return helper(n, d)


# print(Solution().climbStairs(15))
