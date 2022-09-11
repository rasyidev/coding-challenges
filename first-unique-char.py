def first_unique_char(s):
  hmap = {}
  for c in s:
    if c not in hmap:
      hmap[c] = 1
    else:
      hmap[c] +=1
  