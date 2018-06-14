#coding=utf-8

class Solution:
    def containsDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        pass

    def containsDuplicate1(self, nums):

        if len(nums) == len(set(nums)):
            return False
        else:
            return True


    def containsDuplicate2(self, nums):
        nums.sort()

        for i in range(len(nums) - 1):
            if nums[i] == nums[i + 1]:
                return True
        return False

    def containsDuplicate3(self, nums):

        nums_set = set()
        for i in nums:
            if i not in nums_set:
                nums_set.add(i)
            else:
                return True
        return False


import timeit

nums = [1,1,1,3,3,4,3,2,4,2]


def benchmark(nums, fun_nums):
    rep = 10
    r_setup = "from __main__ import Solution, nums"
    for i in range(1, fun_nums + 1):
        r = "Solution().containsDuplicate{:d}(nums)".format(i)
        # print(r)
        # print(rep)
        # print(r_setup)
        print(timeit.repeat(stmt=r, setup=r_setup, repeat=rep, number=1000000))


benchmark(nums, 3)

# [0.6729768584599936, 0.5581572575436079, 0.5722882260891757, 0.5697373115011601, 0.5380829780080902, 0.5714471393371414, 0.5913930237602067, 0.5851391941290638, 0.5614346909371619, 0.5467431295775329]
# [0.8009808596467591, 0.815639235808093, 0.82559400517294, 0.8258002274766323, 0.8576291783226377, 0.8167762239115941, 0.827794499869432, 0.8584098206063491, 0.8045672314339694, 0.8670612808143332]
# [0.48239070909262693, 0.4770486822908371, 0.44168076708375814, 0.47083870452887133, 0.4596533173939026, 0.4346221963939563, 0.4551989946465156, 0.4374378024445633, 0.45209064773951724, 0.4414915324410984]