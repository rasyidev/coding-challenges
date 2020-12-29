'''
Given an array nums of integers, return how many of them contain an even number of digits.
'''

def findNumbers(self, nums):
   """
   :type nums: List[int]
   :rtype: int
   """
   even_count = 0
   for num in nums:
      if(len(str(num)) %2 == 0):
         even_count += 1
   return even_count