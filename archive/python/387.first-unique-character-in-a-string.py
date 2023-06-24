#
# @lc app=leetcode id=387 lang=python3
#
# [387] First Unique Character in a String
#

from collections import Counter


class Solution:
    def firstUniqChar(self, s: str) -> int:
        sm = Counter(s)

        sl = len(s)
        for i in range(sl):
            if sm[s[i]] == 1:
                return i
        return -1
