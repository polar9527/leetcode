#
# @lc app=leetcode id=66 lang=python3
#
# [66] Plus One
#
# https://leetcode.com/problems/plus-one/description/
#
# algorithms
# Easy (40.89%)
# Total Accepted:    368.9K
# Total Submissions: 902K
# Testcase Example:  '[1,2,3]'
#
# Given a non-empty array of digits representing a non-negative integer, plus
# one to the integer.
#
# The digits are stored such that the most significant digit is at the head of
# the list, and each element in the array contain a single digit.
#
# You may assume the integer does not contain any leading zero, except the
# number 0 itself.
#
# Example 1:
#
#
# Input: [1,2,3]
# Output: [1,2,4]
# Explanation: The array represents the integer 123.
#
#
# Example 2:
#
#
# Input: [4,3,2,1]
# Output: [4,3,2,2]
# Explanation: The array represents the integer 4321.
#
#


class Solution:
    def plusOne(self, digits):
        plus = 1
        for i in range(len(digits) - 1, -1, -1):
            digits[i] += plus
            if digits[i] == 10:
                digits[i] = 0
            else:
                plus = 0
                break
        if plus == 1:
            return [1] + digits
        return digits


Solution().plusOne([0])
