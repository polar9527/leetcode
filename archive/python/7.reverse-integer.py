#
# @lc app=leetcode id=7 lang=python3
#
# [7] Reverse Integer
#

import sys


class Solution:
    def reverse(self, x: int) -> int:
        # intMaxSYS = sys.maxsize
        # print(intMaxSYS)

        intMax32 = int((1 << 32)/2) - 1
        # print(type(intMax32))
        # print(intMax32)

        intMin32 = ~intMax32
        # print(type(intMin32))
        # print(intMin32)

        reverse = 0
        pop = 0

        if x == 0:
            return 0

        if x > 0:
            while x != 0:
                pop = x % 10
                x //= 10

                if reverse > intMax32//10 or (reverse == intMax32 and pop > 7):
                    return 0

                reverse = reverse * 10 + pop

            return reverse

        if x < 0:
            while x != 0:
                pop = x % -10
                x = -(x // -10)

                if reverse < -(intMin32//-10) or (reverse == intMin32 and pop < -8):
                    return 0

                reverse = reverse * 10 + pop
            return reverse


if __name__ == "__main__":
    # Solution().reverse(-123)
    Solution().reverse(-2147583647)
    # Solution().reverse(2147583647)
