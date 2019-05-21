#
# @lc app=leetcode id=242 lang=python3
#
# [242] Valid Anagram
#

# from collections import defaultdict


# class Solution:
#     def isAnagram(self, s: str, t: str) -> bool:
#         counter = 0
#         am = defaultdict(int)

#         for char in s:
#             am[char] += 1

#         for char in t:
#             if char not in am:
#                 return False
#             else:
#                 if am[char] == 0:
#                     return False
#                 if am[char] > 0:
#                     am[char] -= 1
#                     if am[char] == 0:
#                         counter += 1
#         return counter == len(am)

from collections import Counter


class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        return Counter(s) == Counter(t)
