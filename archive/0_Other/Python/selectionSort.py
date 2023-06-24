def selectionSort(arr):
    for i in range(len(arr)):
        s = i
        for j in range(i+1, len(arr)):
            if arr[s] > arr[j]:
                s = j
        arr[i], arr[s] = arr[s], arr[i]
    return arr


A = [64, 25, 12, 22, 11]
print(selectionSort(A))
