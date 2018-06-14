# -*- coding: utf-8 -*-

class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        revers_dict = {}

        for index, value in enumerate(nums):
            complement = target - value
            if complement in revers_dict:
                return [revers_dict[complement], index]
            else:
                revers_dict[value] = index
