# [7,1,4,5,3,5,9,2]

def best_time_to_buy_and_sell_stock(prices):
  max_profit = 0 # initial
  l = 0

  # mulai dari elemen ke-1 sebagai pointer right
  for i, price in enumerate(prices[1:]):
    if price < prices[l]:
      l = i
      if price - prices[l] > max_profit:
        max_profit = price - prices[l]
    else:
      if price - prices[l] > max_profit:
        max_profit = price - prices[l]
  return max_profit

print(best_time_to_buy_and_sell_stock([7,1,4,5,3,5,9,2]))

# Input:
# [2,1,4]
# Output:
# 2
# Expected:
# 3