# -*- coding: utf-8 -*-


def qsort(L):
    if len(L) <= 1: return L
    return qsort([lt for lt in L[1:] if lt < L[0]]) + L[0:1] + \
           qsort([ge for ge in L[1:] if ge >= L[0]])


###########  quick_sortA ###########

def partitionA(v, left, right):
    key = v[left]
    low = left
    high = right

    while low < high:
        while (low < high) and (v[high] >= key):
            high -= 1
        v[low] = v[high]
        while (low < high) and (v[low] <= key):
            low += 1
        v[high] = v[low]
        v[low] = key

    return low


def quick_sortA(v, left, right):
    if left < right:
        p = partitionA(v, left, right)
        quick_sortA(v, left, p - 1)
        quick_sortA(v, p + 1, right)
    return v


###########  quick_sortB ###########

def partitionB(array, l, r):
    x = array[r]
    i = l - 1
    for j in range(l, r):
        if array[j] <= x:
            i += 1
            array[i], array[j] = array[j], array[i]
    array[i + 1], array[r] = array[r], array[i + 1]
    return i + 1


def quick_sortB(array, l, r):
    if l < r:
        q = partitionB(array, l, r)
        quick_sortB(array, l, q - 1)
        quick_sortB(array, q + 1, r)


unsorted_list = [3, 14, 2, 12, 9, 33, 99, 35, 3, 14, 2, 12, 9, 33, 99, 35]

# v = quick_sortA(unsorted_list, 0, len(unsorted_list) - 1)
# print(v)

import timeit

q = "qsort(unsorted_list)"
q_setup = "from __main__ import qsort, unsorted_list"

a = "quick_sortA(unsorted_list, 0, len(unsorted_list) - 1)"
a_setup = "from __main__ import quick_sortA, partitionA, unsorted_list"

b = "quick_sortB(unsorted_list, 0, len(unsorted_list) - 1)"
b_setup = "from __main__ import quick_sortB, partitionB, unsorted_list"

rep = 10

t1 = timeit.repeat(stmt=q, setup=q_setup, repeat=rep, number=100000)
print(t1)

t2 = timeit.repeat(stmt=a, setup=a_setup, repeat=rep, number=100000)
print(t2)

t3 = timeit.repeat(stmt=b, setup=b_setup, repeat=rep, number=100000)
print(t3)

# [1.512743847426834, 1.3935659360816517, 1.3717953486743788, 1.4509555515198462, 1.3968778923092744, 1.359945678374844, 1.4597707814130771, 1.5021194769428483, 1.3666104341572751, 1.3644479258518576]
# [1.7739763094115428, 1.68373516603177, 1.6935065673882086, 1.8138007414522583, 1.7247754164013678, 1.6591522238097625, 1.7394750027897459, 1.676691880127752, 1.7191342965065104, 1.7915677492336393]
# [2.768751645584949, 2.8849669570189604, 2.8418294747329327, 2.7589456003349113, 2.8035198925237026, 2.763233785016773, 2.93237840175685, 2.7919729004524143, 2.915946985361316, 2.807815735329399]
