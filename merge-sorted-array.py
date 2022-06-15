# You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

# Merge nums1 and nums2 into a single array sorted in non-decreasing order.

# The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

 

# Example 1:

# Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
# Output: [1,2,2,3,5,6]
# Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
# The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
# Example 2:

# Input: nums1 = [1], m = 1, nums2 = [], n = 0
# Output: [1]
# Explanation: The arrays we are merging are [1] and [].
# The result of the merge is [1].
# Example 3:

# Input: nums1 = [0], m = 0, nums2 = [1], n = 1
# Output: [1]
# Explanation: The arrays we are merging are [] and [1].
# The result of the merge is [1].
# Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.
 

# Constraints:

# nums1.length == m + n
# nums2.length == n
# 0 <= m, n <= 200
# 1 <= m + n <= 200
# -109 <= nums1[i], nums2[j] <= 109
 

# Follow up: Can you come up with an algorithm that runs in O(m + n) time?


def merge_sorted_array(nums1, m, nums2, n):
  p1 = m - 1 # pointer 1 di elemen terakhir nums1 sebelum nol
  p2 = n - 1 # pointer 2 di elemen terakhir nums2

  # m = 0 => return nums2 | n = 0 => return nums1
  if m == 0:
    return nums2
  if n == 0:
    return nums1

  for i in range(5, 1, -1): # loop index 5 -> 2 dari nums1
    print(nums1)
    print(nums2[p2], "vs", nums1[p1])
    if nums2[p2] >= nums1[p1]:
      nums1[i] = nums2[p2]
      p2 -= 1 # geser pointer ke kiri

    else:
      print(nums2[p2], "vs", nums1[p1])
      nums1[i] = nums1[p1]
      p1 -=1  # geser pointer ke kiri
  return nums1

# Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
# print(merge_sorted_array([1,2,3,0,0,0], 3, [2,5,6], 3)) pass

print(merge_sorted_array([0], 0, [1], 1))