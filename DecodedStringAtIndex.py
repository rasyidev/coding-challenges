def decodeAtIndex(S, K):
    res = ""
    for char in S:
        if(char.isdigit() == False): # which means it is not a number
            res += char
        else:
            res = res * (int(char))
    return length(res)
        
res = decodeAtIndex("a2345678999999999999999", 10)
print(res)