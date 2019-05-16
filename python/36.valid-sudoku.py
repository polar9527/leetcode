#
# @lc app=leetcode id=36 lang=python3
#
# [36] Valid Sudoku
#
import collections


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = collections.defaultdict(set)
        columns = collections.defaultdict(set)
        boxs = collections.defaultdict(set)

        for i_row, row_cells in enumerate(board):
            for i_column, cell in enumerate(row_cells):
                if cell != ".":
                    if cell in rows[i_row] or cell in columns[i_column] or cell in boxs[(i_row//3, i_column//3)]:
                        return False
                    rows[i_row].add(cell)
                    columns[i_column].add(cell)
                    boxs[(i_row//3, i_column//3)].add(cell)
        return True
