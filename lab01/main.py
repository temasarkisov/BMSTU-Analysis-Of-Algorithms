def levenshtein_distance(s1, s2):
    n, m = len(s1), len(s2)
    if n > m:
        s1, s2 = s2, s1
        n, m = m, n

    current_row = range(n + 1)
    for i in range(1, m + 1):
        previous_row, current_row = current_row, [i] + [0] * n
        for j in range(1, n + 1):
            add, delete, substitution = previous_row[j] + 1, current_row[j - 1] + 1, previous_row[j - 1]
            if s1[j - 1] != s2[i - 1]:
                substitution += 1
            current_row[j] = min(add, delete, substitution)

    return current_row[n]

def damerau_levenshtein_distance(s1, s2):
    d = {}
    for i in range(len(s1) + 1):
        d[(i, 0)] = i
    for j in range(len(s2) + 1):
        d[(0, j)] = j

    for i in range(1, len(s1) + 1):
        for j in range(1, len(s2) + 1):
            if s1[i - 1] == s2[j - 1]: 
                cost = 0
            else: 
                cost = 1
            d[(i, j)] = min(d[(i - 1, j)] + 1, d[(i , j - 1)] + 1, d[(i - 1, j - 1)] + cost)
            if i > 1 and j > 1 and s1[i - 1]==s2[j - 2] and s1[i - 2] == s2[j - 1]:
                d[(i, j)] = min(d[(i, j)], d[i - 2,j - 2] + cost) 
 

    #for i in range(len(s1) + 1):
    #    print()
    #    for j in range(len(s2) + 1):
    #        print(d[(i, j)], end=' ')
    #print("\n")

    return d[len(s1), len(s2)]


s1 = "abcd"
s2 = "adb"
print("Levenshtein distance of '%s' and '%s' is %d."%(s1, s2, levenshtein_distance(s1, s2)))   
print("Damerau - Levenshtein distance of '%s' and '%s' is %d."%(s1, s2, damerau_levenshtein_distance(s1, s2)))   