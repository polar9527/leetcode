# -*- coding: utf8 -*-

"""
约瑟夫环问题起源于一个犹太故事。约瑟夫环问题的大意如下

罗马人攻占了桥塔帕特，41个人藏在一个山洞中躲过了这场浩劫。这41个人中，包括历史学家约瑟夫和他的一个朋友。剩余的39个人为了表示不向罗马人屈服，
决定集体自杀。大家决定了一个自杀方案，所有这41个人围成一个圆圈，由第1个人开始顺时针报数，每报数为3的人就立刻自杀，然后再由下一个人继续报数，
仍然是每报数为3的人就立刻自杀……，直到所有人都自杀身亡为止。

问题变为（N-1）个人，报到为（M-1）的人自杀，问题规模减小了。这样一直进行下去，最后剩下编号为0的人。用函数表示：

F(1)=0

当有2个人的时候（N=2），报道（M-1）的人自杀，最后自杀的人是谁？应该是在只有一个人时，报数时得到的最后自杀的序号加上M，因为报到M-1的人已经自杀，
只剩下2个人，另一个自杀者就是最后自杀者，用函数表示：

F(2)=F(1)+M

可以得到递推公式：

F(i)=F(i-1)+M

因为可能会超出总人数范围，所以要求模

F(i)=(F(i-1)+M)%i
"""
def josephus1(n ,k):
    link = list(range(1 , n +1))
    ind =0
    for loop_i in range( n -1):
        ind = (ind + k)% len(link)
        ind -= 1
        print('Kill:' ,link[ind])
        del link[ind]
        if ind ==-1: # the last element of link
            ind =0
    print('survice :' ,link[0])

def josephus2(n ,k):
    s = 0
    for i in range(1, n):
        s = (s + k) % i

    index = s + 1
    return index

def josephus3(n ,k):

    if n > 1:
        ret =  (josephus3(n - 1, k) + k) % n
    else:
        ret =  0

    index = ret + 1
    return index


if __name__ == '__main__':

    # josephus(100000 ,300)
    # print('- ' *30)
    # josephus(10 ,5)
    print('- ' *30)
    ret = josephus3(41 ,3)
    print(ret)
