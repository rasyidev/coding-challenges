def findMaxConsecutiveOnes(nums):
    consecutive_count = 0
    # print("len nums:", len(nums))
    temp = []

    for i in range(len(nums)):
        if(nums[i] == 1):
            consecutive_count +=1
            temp.append(consecutive_count)
        else:
            consecutive_count = 0
    return max(temp)
res = findMaxConsecutiveOnes([1,0,1,0,1,1,1,1,0,1,1,1,1,1])
print(res)
        