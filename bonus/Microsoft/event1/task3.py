# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")


def solution(A, B, S):
    # write your code in Python (Python 3.6)

    def dfs(choosed, slots, i):
        if i > len(A) - 1:
            return True
        # we could choose A[i] or B[i]
        if slots[choosed] == True:
            return False
        slots[choosed] = True

        # check index out of range
        if i > len(A) - 2:
            return True

        return dfs(A[i + 1], slots, i + 1) or dfs(B[i + 1], slots, i + 1)

    slots = {}
    for i in range(S + 1):
        slots[i] = False

    return dfs(B[0], slots, 0) or dfs(A[0], slots, 0)


print(solution([1, 1, 3], [2, 2, 1], 3))
print(solution([3, 2, 3, 1], [1, 3, 1, 2], 3))
print(solution([2, 5, 6, 5], [5, 4, 2, 2], 8))
print(solution([1, 2, 1, 6, 8, 7, 8], [2, 3, 4, 7, 7, 8, 7], 10))
