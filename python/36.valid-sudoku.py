#
# @lc app=leetcode id=36 lang=python3
#
# [36] Valid Sudoku
#

import collections


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        row = collections.defaultdict(set)
        column = collections.defaultdict(set)
        square = collections.defaultdict(set)

        for iRow in range(10):
            print(iRow)
