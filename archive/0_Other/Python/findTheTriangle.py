# 找“三角”
# 找到输入的二元组数组中，所有可以组成三元组的组合并输出一个数组
# 例子：
# 输入[(1, 2), (2, 3), (1, 3), (1, 4), (4, 5), (3, 4)]
# 输出[(1, 2, 3), (1, 3, 4)]
# 输出的每个三元组中，从里面任意取两个元素组成一个二元tuple，都能在输入中找到


def ft(arrT2):
    d = {}
    arrS3 = []
    arrT3 = []

    for t in arrT2:
        if t[0] in d.keys():
            d[t[0]].add(t[1])
        else:
            d[t[0]] = set([t[1], ])
        if t[1] in d.keys():
            d[t[1]].add(t[0])
        else:
            d[t[1]] = set([t[0], ])

    for intKey, pairSet in d.items():
        for x in pairSet:
            if x in d.keys():
                for y in d[x]:
                    if y in pairSet:
                        tempSet = set([intKey, x, y])
                        if tempSet not in arrS3:
                            arrS3.append(set([intKey, x, y]))
    for s in arrS3:
        arrT3.append(tuple(s))

    return arrT3


inputData = [(1, 2), (2, 3), (1, 3), (1, 4), (4, 5), (3, 4)]
outputData = [(1, 2, 3), (1, 3, 4)]

print(ft(inputData))
