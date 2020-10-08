def get_three(n):
    T = "012"
    q, r = divmod(n, 3)
    if q == 0:
        return T[r]
    else:
        return get_three(q) + T[r]

def solution(n):
    answer = 0
    convert = get_three(n)[::-1]
    return int(convert, 3)