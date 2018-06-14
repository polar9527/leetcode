# coding=utf8

from functools import reduce


class Solution:
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        return reduce(lambda x, y: x^y, nums)

    def singleNumbers2(self, nums):
        single_nums_set = set()

        single_nums_set.add(nums[0])

        for i in range(1, len(nums)):
            if nums[i] not in single_nums_set:
                single_nums_set.add(nums[i])
            else:
                single_nums_set.remove(nums[i])
        return single_nums_set.pop()

nums =  [4,1,2,1,2]
ret = Solution().singleNumbers2(nums)
print(ret)
