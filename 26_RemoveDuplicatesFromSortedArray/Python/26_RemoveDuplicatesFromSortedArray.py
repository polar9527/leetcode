# -*- coding: utf-8 -*-

import unittest
import timeit


def removeDuplicates(nums):
    if len(nums) == 0:
        return 0

    slower = 0

    for faster in range(1, len(nums)):
        if nums[slower] != nums[faster]:
            slower += 1
            nums[slower] = nums[faster]

    return slower + 1


def removeDuplicatesBuildin(nums):
    nums[:] = sorted(list(set(nums)))
    return len(nums)


class TestRemoveDuplicates(unittest.TestCase):

    def setUp(self):
        self.nums = [1, 1, 3, 4, 4, 5, 5, 5, 6, 7, 7, 7]
        self.result = 6

    def test_removeDuplicates(self):
        result = removeDuplicates(self.nums)
        self.assertEqual(self.result, result)


class TestRemoveDuplicatesBuildin(unittest.TestCase):

    def setUp(self):
        self.nums = [1, 1, 3, 4, 4, 5, 5, 5, 6, 7, 7, 7]
        self.result = 6

    def test_removeDuplicatesBuildin(self):
        result = removeDuplicatesBuildin(self.nums)
        self.assertEqual(self.result, result)


if __name__ == '__main__':
    unittest.main(verbosity=2)
    # rep = 8
    # print(timeit.timeit("testA()", setup="from __main__ import testA", number=1))

    # print(timeit.repeat("removeDuplicates(nums)", setup="nums = [1,1,3,4,4,5,5,5,6,7,7,7];"
    #                     "from __main__ import removeDuplicates",
    #                     repeat=rep))

    # print(timeit.repeat("removeDuplicatesBuildin(nums)", setup="nums = [1,1,3,4,4,5,5,5,6,7,7,7];"
    #                     "from __main__ import removeDuplicatesBuildin",
    #                     repeat=rep))

    # [1.7508393086419753, 1.7301957530864196, 1.7445479506172838, 1.711339851851852, 2.061765135802469, 1.8548669629629622, 1.7146026666666678, 1.7883930864197524]
    # [1.2516025679012355, 1.2129351111111113, 1.1901202962962962, 1.2308744691358022, 1.2144853333333323, 1.1707535802469131, 1.2358640987654326, 1.150896592592595]
