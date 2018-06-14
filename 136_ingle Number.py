# coding=utf8

from functools import reduce


class Solution:
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        return reduce(lambda x, y: x^y, nums)

nums =  [4,1,2,1,2]
ret = Solution().singleNumber(nums)
print(ret)
