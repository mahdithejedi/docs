## All implimentaion of QUICKSORT

# CLRS

def partition(arr, first, end):
    pivot = arr[end]
    i = first -1 
    for j in range(first, end):
        if arr[j] <= pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]
    arr[i+1],arr[end] = arr[end], arr[i+1]
    return i + 1

def quick_sort(arr, first, end):
    """
        CLRS implementation
    """
    if first < end:
        # Prtiion will set pivot and 
        # Smaller than pivot go to left
        # Bigger than pivot go to right
        q = partition(arr, first, end)
        quick_sort(arr,first, q-1)
        quick_sort(arr,q+1, end)