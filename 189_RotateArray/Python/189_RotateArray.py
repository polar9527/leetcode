# -*- coding: utf-8 -*-

class Solution:
    def rotate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        pass

    def rotate1(self, nums, k):
        n = len(nums)
        k = k % n
        nums[:] = nums[n - k:] + nums[:n - k]

    def rotate2(self, nums, k):
        nums.reverse()

        k = k % len(nums)

        start = 0
        end = k - 1
        while start < end:
            nums[start], nums[end] = nums[end], nums[start]
            start += 1
            end -= 1

        start = k
        end = len(nums) - 1
        while start < end:
            nums[start], nums[end] = nums[end], nums[start]
            start += 1
            end -= 1

    def rotate3(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        k = k % n
        for i in range(k):
            val = nums.pop()
            nums.insert(0, val)

    def rotate4(self, nums, k):
        nums[:] = reversed(nums)

        nums[:k] = reversed(nums[:k])

        nums[k:] = reversed(nums[k:])


import timeit

k = 3
nums = [1, 2, 3, 4, 5, 6, 7]


def benchmark(k, nums, fun_nums):
    rep = 10
    r_setup = "from __main__ import Solution, k, nums"
    for i in range(1, fun_nums + 1):
        r = "Solution().rotate{:d}(nums, k)".format(i)
        # print(r)
        # print(rep)
        # print(r_setup)
        print(timeit.repeat(stmt=r, setup=r_setup, repeat=rep, number=100000))


benchmark(k, nums, 4)

# [0.07254606128915905, 0.0740433470143372, 0.07518468170036349, 0.07318013611472923, 0.07173618378609437, 0.0788137230384866, 0.07332749431406732, 0.07241015989352284, 0.07251998718418229, 0.08082419455707013]
# [0.11340773934744408, 0.11103262542138326, 0.10821148627534172, 0.1206753035164092, 0.11598868067794244, 0.10949346310336372, 0.11068102006639413, 0.10819923934724662, 0.1106474397796815, 0.1085642768169206]
# [0.12815778152033364, 0.132805688263532, 0.13213961340003566, 0.1459940495731944, 0.13685428565446234, 0.12985615390358962, 0.13381704748687095, 0.13084262420854254, 0.13398652916921971, 0.12944765959228777]
# [0.15899633658825074, 0.1822615493470212, 0.16870617525670184, 0.16197708086172558, 0.16139238881073226, 0.1597070534799645, 0.15619416042764733, 0.16097954881526721, 0.17717907418754475, 0.164320589812057]
