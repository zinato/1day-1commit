def solution(n, p, c):
    days = 0
    cost = 0
    rest = 0
    is_half = False
    is_cancel = False

    for _p, _c in zip(p, c):
        days += 1

        if is_cancel:
            break

        rest += p

        if rest >= _c:
            rest -= _c
            unit = 50 if is_half else 100
            cost += (unit * _c)
            is_half = False
        else:
            if is_half:
                is_cancel = True
            else:
                is_half = True

    answer = "%.2f" % (cost/days)
    print(answer)