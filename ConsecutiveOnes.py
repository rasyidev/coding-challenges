def findMaxConsecutiveOnes(nums):
    consecutive_count = 0
    # print("len nums:", len(nums))
    temp = 0

    for i in range(len(nums)):
        if nums[i] == 1:
            consecutive_count +=1
            if temp > consecutive_count:
                temp = consecutive_count
        else:
            consecutive_count = 0
    return temp
res = findMaxConsecutiveOnes([1,0,1,0,1,1,1,1,0,1,1,1,1,1])
print(res)
        