# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")


def solution(S):
    # find the middle R
    index_arr = []
    for i, s in enumerate(S):
        if s == "R":
            index_arr.append(i)

    if len(index_arr) == 0:
        return 0

    mid = len(index_arr) // 2

    # loop from the middle R to the fitrst R
    print(index_arr)

    left_shift = 0
    left_mid = index_arr[mid]
    for i in range(mid - 1, -1, -1):
        # the shift when i == 1 should be left_mid - index_arr[i]
        left_shift += left_mid - 1 - index_arr[i]
        left_mid -= 1
        if left_shift > 1000000000:
            return -1

    right_shift = 0
    right_mid = index_arr[mid]
    for i in range(mid + 1, len(index_arr)):
        right_shift += index_arr[i] - right_mid - 1
        right_mid += 1
        if right_shift > 1000000000:
            return -1

    print(left_shift, right_shift)

    return left_shift + right_shift


solution("WRRWWR")
solution("WWRWWWRWR")
solution("WWW")
solution("RW")
