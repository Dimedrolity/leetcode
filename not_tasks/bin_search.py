def binarySearch(nums, target):
    low = 0
    high = len(nums) - 1
    middle = 0

    while low <= high:
        middle = (high + low) // 2

        # target equals nums[middle], return middle (index of target)
        if target == nums[middle]:
            return middle
        # Take the upper half
        elif target > nums[middle]:
            low = middle + 1

        # Take the lower half
        else:
            high = middle - 1

    # If we reach this line, target is not present in nums
    return -1

def binarySearchRec(nums, target, low, high):
    if low <= high:
        middle = (high + low) // 2

        # target equals nums[middle], return middle (index of target)
        if target == nums[middle]:
            return middle

        # Take the upper half
        elif target > nums[middle]:
            return binarySearchRec(nums, target, middle + 1, high)

        # Take the lower half
        else:
            return binarySearchRec(nums, target, low, middle - 1)

    # If we reach this line, target is not present in nums
    else:
        return -1