def isAnagram(self, s: str, t: str) -> bool:
    letters = {}
    for l in s:
        if l not in letters:
            letters[l] = 0
        letters[l] += 1

    for i in t:
        if i in letters:
            letters[i] -= 1
            if letters[i] == 0:
                del letters[i]
        else:
            return False
    return len(letters) == 0