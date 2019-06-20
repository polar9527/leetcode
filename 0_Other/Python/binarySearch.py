def binarySearch(arr, l, r, x):

    while l < r:
        m = (l + r)//2
        if arr[m] < x:
            l = m + 1
        elif arr[m] > x:
            r = m
        else:
            return m

    return -1


# 测试数组
arr = [2, 3, 4, 10, 40]
x = 10

# 函数调用
result = binarySearch(arr, 0, len(arr)-1, x)

if result != -1:
    print("元素在数组中的索引为 %d" % result)
else:
    print("元素不在数组中")
