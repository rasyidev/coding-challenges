'''
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
A subarray is a contiguous part of an array.

 
Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Example 2:

Input: nums = [1]
Output: 1
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
 

Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104

Example : [-2,1,-3,4,-1,2,1,-5,4]

step 0 : currSum = maxResult =-2
step 1 : currSum = max(-2+1,1) =1  --> maxResult = 1
step 2 : currSum = max(1+-3,-3) = -2  --> maxResult = 1
step 3 : currSum = max(-2 + 4,4) = 4  --> maxResult = 4
step 4 : currSum = max( 4 - 1 ,-1) = 3  --> maxResult = 4
step 5 : currSum = max( 3 + 2 ,2) = 5  --> maxResult = 5
step 5 : currSum = max( 5 + 1 ,1) = 6 --> maxResult = 6
step 6 : ...........
step 7 : ...........
'''

def maximum_subarray(nums):
  max_sum = nums[0]
  max_contaguous = max_sum
  for i in range(len(nums)):
    if i != 0:
      max_contaguous = max(nums[i], max_contaguous + nums[i])
      if max_contaguous > max_sum:
        max_sum = max_contaguous

  return max_sum 


nums = [-2,1,-3,4,-1,2,1,-5,4]
print(maximum_subarray(nums))
