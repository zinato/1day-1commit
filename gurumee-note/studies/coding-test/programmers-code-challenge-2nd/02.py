def get_lt(arr, n):
    res = []
    f = n // 2
    
    for i in range(f):
        res.append(arr[i][:f])
    
    return res

def get_rt(arr, n):
    res = []
    f = n // 2
    
    for i in range(f, n):
        res.append(arr[i-f][f:])
    
    return res

def get_lb(arr, n):
    res = []
    f = n // 2
    
    for i in range(f):
        res.append(arr[i+f][:f])
    
    return res

def get_rb(arr, n):
    res = []
    f = n // 2
    
    for i in range(f, n):
        res.append(arr[i][f:])
    
    return res

def compress(arr, n):
    if n == 1:
        return arr[0][0]
    
    lt = compress(get_lt(arr, n), n // 2)
    rt = compress(get_rt(arr, n), n // 2)
    lb = compress(get_lb(arr, n), n // 2)
    rb = compress(get_rb(arr, n), n // 2)
    
    if lt is rt and rt is lb and lb is rb:
        return lt
    
    return [lt, rt, lb, rb]

def count(qt, x):
    if type(qt) == int:
        if qt == x:
            return 1
        else:
            return 0
    
    lt = count(qt[0], x)
    rt = count(qt[1], x)
    lb = count(qt[2], x)
    rb = count(qt[3], x)
    return lt + rt + lb + rb


def solution(arr):
    answer = []
    n = len(arr)
    qt = compress(arr, n)
    print(qt)
    answer.append(count(qt, 0))
    answer.append(count(qt, 1))
    return answer


if __name__ == '__main__':
    arr = [[1, 0, 0, 0] * 2] * 8
    answer = solution(arr)
    print(answer)

    a1 = [1, 0 ,0]
    b1 = [1, 0, 0]
    print(a1 is b1)