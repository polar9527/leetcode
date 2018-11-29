# -*- coding: utf-8 -*-

"""
给定n个字符串,每个字符串的组成字母都属于A-J,然后我们需要确定一个映射使得A-J对应0-9并且所有字符串对应的整数加起来最大。另外每一个字符串都
不允许有前置0。
题目比较简单,我们分析每一个字符串中每一个字母的权重(1,10,100…),然后将所有字母的权重加起来,
按照权重进行排序这样就可以获得最大映射的序列。但是由于题目要求不能有前置0,所以如果出现的字母种类达到10种,
那么就需要找到一个最小权重且不为首字母的字母,将它提前赋值为0。然后剩下的字母从9开始倒序赋值即可。

输入例子:
2
ABC
BCA
输出例子:
1875

"""
from functools import reduce


def solution(strings):
    def toint(ch):return ord(ch)-65

    power = [10**i for i in range(10)]
    count = {}
    firsts = set()
    for string in strings:
        firsts.add(toint(string[0]))
        val = 1
        for index in [toint(ch) for ch in string[::-1]]:
            if index in count:
                count[index]+= val
            else:
                count[index] = val
            val*=10

    count = sorted(count.items(),key=lambda x:x[1])
    result = [0]*10
    print(len(count))
    if len(count)==10:
        for i,index in count:
            if i not in firsts:
                count.remove((i,index))
                result[i] = 0
                break

    init=9
    for i,j in count[::-1]:
        result[i],init = init,init-1
    ans = 0
    print(result,count,firsts)
    for string in strings:
        ans += reduce(lambda x, y:x+result[toint(y[0])]*y[1], zip(string[::-1],power), 0)
    return ans


strings = ['ABC', 'BCA']

ret = solution(strings)

print(ret)