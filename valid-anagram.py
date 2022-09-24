def isAnagram(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False

    hashMap = {}
    for i in range(len(s)):
        if s[i] in hashMap.keys():
            hashMap[s[i]] += 1
        else:
            hashMap[s[i]] = 1

        if t[i] in hashMap.keys():
            hashMap[t[i]] -= 1
        else:
            hashMap[t[i]] = -1

    for v in hashMap.values():
        if v != 0:
            return False

    return True

print(isAnagram("anagram", "nagaram"))