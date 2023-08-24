class Segment:
    def __init__(self):
        self.start = 0
        self.end = 0
        self.value = 0


def solution(S, B):
    # write your code in Python (Python 3.6)

    # Golang to Python
    pathholes = []
    seg = Segment()
    for i in range(len(S)):
        print(seg.start, seg.end, seg.value)
        if S[i] == ".":
            # check if seg is empty
            if seg.start != seg.end:
                # put it into map
                seg.value = seg.end - seg.start
                pathholes.append(seg)
            seg = Segment()
            continue
        if S[i] == "X":
            # check if we are in a segement
            if seg.start == seg.end:
                # create a Segement
                seg = Segment()
                seg.start = i
                seg.end = i + 1
            else:
                # put current segement in the map
                seg.end = i + 1

    # if segement is not empty put it in the map
    if seg.start != seg.end:
        seg.value = seg.end - seg.start
        pathholes.append(seg)
    # start from max of pathholes
    pathholes.sort(key=lambda x: x.value, reverse=True)
    fixed = 0
    # fix
    for i in range(len(pathholes)):
        if pathholes[i].value > B:
            # fix as many as possible
            fixed += B
            break
        else:
            B -= pathholes[i].value + 1
            fixed += pathholes[i].value

    return fixed


print(solution("...XXX..X....XXX", 7))
print(solution("........X....XXX", 4))
print(solution("X.X.XXX...X", 14))
print(solution("..", 5))
